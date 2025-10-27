package resources

import (
	"errors"
	"log"
	"strings"

	"github.com/studio-b12/gowebdav"
)

func searchRecursive(client *gowebdav.Client, path string, searchText string) (string, error) {
	// Ne pas lister la racine elle-même, seulement son contenu
	if path == "/" {
		path = ""
	}

	files, err := client.ReadDir(path)
	if err != nil {
		// On ignore les erreurs "Not Implemented" (dossiers non-listables) ou "Not Found"
		if strings.Contains(err.Error(), "501") || strings.Contains(err.Error(), "404") {
			return "", nil
		}
		log.Printf("Erreur ReadDir pendant la recherche : %v (path: %s)", err, path)
		return "", err
	}

	for _, file := range files {
		fullPath := path + "/" + file.Name()

		// Vérifie si le nom du fichier correspond (insensible à la casse)
		if !file.IsDir() && strings.Contains(strings.ToLower(file.Name()), strings.ToLower(searchText)) {
			log.Printf("Correspondance trouvée : %s", fullPath)
			return fullPath, nil // On retourne le premier match
		}

		// Si c'est un dossier, on continue la recherche à l'intérieur
		if file.IsDir() {
			// Évite de descendre dans des dossiers "inutiles"
			if file.Name() == "." || file.Name() == ".." {
				continue
			}

			// Appel récursif
			foundPath, err := searchRecursive(client, fullPath, searchText)
			if err != nil {
				continue // Continue la recherche même si un sous-dossier échoue
			}
			if foundPath != "" {
				return foundPath, nil // Match trouvé dans un sous-dossier
			}
		}
	}

	return "", nil // Pas de match dans ce dossier
}

// GetMediaPath recherche un fichier sur le WebDAV en fonction d'un texte
// et renvoie le chemin complet du premier fichier correspondant.
func GetMediaPath(client *gowebdav.Client, searchText string) (string, error) {
	if searchText == "" {
		return "", errors.New("le texte de recherche ne peut pas être vide")
	}

	// ATTENTION: Démarrer la recherche depuis la racine "/"
	// peut être extrêmement lent et causer des crashs mémoire.
	// Il est RECOMMANDÉ de démarrer depuis un dossier spécifique,
	// par exemple "/movies" ou "/series".
	log.Printf("Démarrage de la recherche de média pour : %s", searchText)

	// Pour de meilleures performances, vous devriez lancer la recherche
	// depuis des dossiers racines spécifiques, ex:
	// rootFolders := []string{"/movies", "/series", "/anime"}
	// for _, folder := range rootFolders {
	//    foundPath, err := searchRecursive(client, folder, searchText)
	//    ... (et retourner si trouvé)
	// }

	// Pour l'instant, on commence à la racine.
	foundPath, err := searchRecursive(client, "/", searchText)
	if err != nil {
		return "", err
	}

	if foundPath == "" {
		return "", errors.New("aucun média trouvé pour: " + searchText)
	}

	return foundPath, nil
}
