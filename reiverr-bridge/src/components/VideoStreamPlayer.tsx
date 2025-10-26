import { useState, useRef, useEffect } from 'react';

// L'URL de votre API qui renvoie le lien du stream
// Pour cet exemple, on va simuler la réponse de l'API.
// const API_URL = 'https://api.example.com/get-video-stream';
const MOCK_VIDEO_URL = 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4';

function VideoStreamPlayer() {
  // 'videoSrc' gardera en mémoire l'URL du stream
  const [videoSrc, setVideoSrc] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  
  // On utilise useRef pour contrôler le lecteur vidéo (ex: .play())
  const videoRef = useRef<HTMLVideoElement>(null);
  // Fonction appelée au clic sur le bouton
  const handlePlayClick = () => {
    setIsLoading(true);
    setError(null);
    setVideoSrc(''); // Réinitialise la vidéo précédente

    // --- Simulation d'un appel API ---
    // Dans un cas réel, vous utiliseriez fetch() ou axios()
    // fetch(API_URL)
    //   .then(response => response.json())
    //   .then(data => {
    //     setVideoSrc(data.streamUrl); // Supposons que l'API renvoie { streamUrl: "..." }
    //     setIsLoading(false);
    //   })
    //   .catch(err => {
    //     setError('Impossible de charger la vidéo.');
    //     setIsLoading(false);
    //   });

    // --- Début de la simulation ---
    console.log('Appel API simulé pour récupérer le stream...');
    setTimeout(() => {
      console.log('Stream reçu !');
      setVideoSrc(MOCK_VIDEO_URL);
      setIsLoading(false);
    }, 1500); // Simule un délai réseau de 1.5s
    // --- Fin de la simulation ---
  };

  // Cet effet se déclenche quand 'videoSrc' change
  useEffect(() => {
    // Si on a une URL et que le lecteur vidéo est prêt
    if (videoRef.current && videoSrc) {
      // On charge la nouvelle source et on lance la lecture
      videoRef.current.load();
      videoRef.current.play().catch(err => {
        console.error("Erreur lors de la lecture auto :", err);
        // La lecture auto est souvent bloquée par les navigateurs
        // Il faut afficher les contrôles pour que l'utilisateur clique
      });
    }
  }, [videoSrc]); // Dépendance : se relance si 'videoSrc' est modifié

  return (
    <div>
      <h2>Lecteur de Stream</h2>
      
      <button onClick={handlePlayClick} disabled={isLoading}>
        {isLoading ? 'Chargement...' : '▶️ Lancer le Stream (Appel API)'}
      </button>

      {error && <p style={{ color: 'red' }}>{error}</p>}

      {/* Le lecteur vidéo s'affiche seulement si 'videoSrc' n'est pas null.
        On ajoute 'controls' pour que l'utilisateur puisse gérer la lecture.
      */}
      {videoSrc && (
        <div style={{ marginTop: '20px' }}>
          <video 
            ref={videoRef}
            width="600" 
            controls 
            autoPlay // 'autoPlay' peut être bloqué par le navigateur
          >
            <source src={videoSrc} type="video/mp4" />
            Votre navigateur ne supporte pas l'élément vidéo.
          </video>
        </div>
      )}
    </div>
  );
}

export default VideoStreamPlayer;