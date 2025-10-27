package resources

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/studio-b12/gowebdav"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type streamRequest struct {
	Token string `json:"token"`
}

func SocketStream(c *gin.Context, client *gowebdav.Client, cache *CacheToken) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Upgrade error to websocket: %v", err)
		return
	}
	// S'assure de fermer la connexion à la fin
	defer conn.Close()

	// 2. Boucle infinie pour écouter les messages sur ce WebSocket
	// (Permet au client de demander plusieurs fichiers sans se reconnecter)
	for {
		var req streamRequest

		// 3. Attend un message JSON du client (ex: {"path": "/films/mon_film.mkv"})
		err := conn.ReadJSON(&req)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Erreur de lecture WebSocket: %v", err)
			}
			// Le client s'est déconnecté, on arrête la boucle
			break
		}

		cache.cacheMutex.RLock() // Verrouille la map en lecture
		path, ok := cache.MediaCache[req.Token]
		cache.cacheMutex.RUnlock() // Déverrouille

		if !ok {
			log.Printf("Token invalide ou expiré reçu : %s", req.Token)
			conn.WriteJSON(gin.H{"error": "Token invalide ou expiré"})
			continue // Attend un nouveau message
		}
		// (Pour plus de sécurité, vous pourriez supprimer le token du cache ici)
		cache.cacheMutex.Lock()
		delete(cache.MediaCache, req.Token)
		cache.cacheMutex.Unlock()
		// --- FIN LOGIQUE DE TOKEN ---

		log.Printf("Requête de stream WS reçue pour le token %s (path: %s)", req.Token, path)

		// 4. Récupère le flux de données depuis le WebDAV
		stream, err := streamFromFile(client, path)
		if err != nil {
			log.Printf("Erreur WebDAV (StreamFile): %v", err)
			// Envoie un message d'erreur JSON au client
			conn.WriteJSON(gin.H{"error": err.Error()})
			continue // Attend la prochaine requête
		}
		// 'stream' est un io.ReadCloser

		buf := make([]byte, 32*1024)
		for {
			n, err := stream.Read(buf)
			if n > 0 {
				if wErr := conn.WriteMessage(websocket.BinaryMessage, buf[:n]); wErr != nil {
					log.Printf("Erreur d'écriture WebSocket: %v", wErr)
					break
				}
			}

			if err == io.EOF {
				log.Printf("Stream terminé pour : %s", path)
				conn.WriteMessage(websocket.TextMessage, []byte("STREAM_END"))
				break
			}

			if err != nil {
				log.Printf("Erreur de lecture du stream: %v", err)
				conn.WriteJSON(gin.H{"error": "Erreur pendant la lecture du stream"})
				break
			}
		}

		stream.Close()
	}
}
