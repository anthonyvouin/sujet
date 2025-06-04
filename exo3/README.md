# WORD GUESS (11 points)

Le programme sera un exécutable codé en un package unique dans le dossier actuel.
Seules les librairies standards du SDK Golang sont autorisées.

## Spécifications
Ecrire un programme Golang qui permet de jouer à deviner un mot secret.
Au lancement, un mot secret d'une longueur précise est choisi (ou généré).
On peut ensuite faire des essais pour deviner le mot secret.

## Algo du jeu
À chaque essai, le jeu indique le statut pour chaque lettre du mot essayé :
- Lettre présente à cette position dans le mot secret
- Lettre présente à une position différente dans le mot secret
- Lettre absente du mot secret

## Configuration
Les paramètres du jeu sont les suivants :
- longueur du mot à deviner
- nombre d'essais maximum
- méthode pour générer le mot secret
- méthode pour deviner le mot secret

## Méthodes de génération du mot secret 
Implémenter au moins 2 méthodes pour générer le mot secret :
- Méthode 1 : mot choisi à partir d'un dictionnaire
- Méthode 2 : suite de lettres aléatoires

## Méthodes de jeu
Implémenter au moins 2 méthodes pour deviner le mot :
- Méthode 1 : utilisateur humain.
L'utilisateur saisi des mots dans sur l'entrée standard.
Le résultat de chaque essai s'affiche sur la sortie standard.

- Méthode 2 : utilisateur IA.
Le jeu est automatisé, le programme joue à deviner le mot. 

## Tests
Ecrivez les tests unitaires pour l'algorithme principal du jeu.
