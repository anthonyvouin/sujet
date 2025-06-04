# Gestion d'une liste d'étudiants (5 points)

Tout le code de cet exercice est à implémenter dans le fichier `students.go`.
Le fichier `students_test.go` ne doit PAS être modifié.
Le code implémenté doit faire passer avec succès les tests unitaires du fichier `students_test.go` avec une couverture de code de 100%.
Seules les librairies standards du SDK Golang sont autorisées.

## Instructions

### 1. Structure Student
Définir une structure `Student` qui contient les champs suivants :
- `Name` (chaîne de caractères)
- `Age` (nombre entier)
- `Grade` (nombre à virgule 64 bits)

### 2. Structure StudentList
Définir une structure `StudentList` qui contient un champ `students` de type slice de `Student`.

### 3. Fonction NewStudent
Implémenter la fonction `NewStudent` :
- Paramètres : `name` (chaîne de caractères), `age` (entier) et `grade` (nombre à virgule)
- Crée et retourne un pointeur de `Student` avec une erreur éventuelle
- Contraintes : `name` ne doit pas être vide, `age` doit être compris entre 1 et 99, et `grade` doit être compris entre 0 et 20.
- Si l'une des contraintes n'est pas respectée, renvoyer nil et une erreur.

### 4. Méthodes de StudentList
Implémenter les méthodes suivantes pour la structure `StudentList` :
- `AddStudents` : accepte un nombre quelconque de `Student` en paramètre et les ajoute à la liste
- `RemoveStudent` : accepte un paramètre `name` (chaîne de caractères), et retire de la liste les étudiants dont le nom correspond
- `Sort` : trie les étudiants par ordre de note décroissante et renvoie une nouvelle `StudentList`
- `Print` : accepte un paramètre `out` de type `io.Writer`, écrit sur `out` la liste des étudiants (une ligne par étudiant) avec le format suivant : "Name (age): grade".

