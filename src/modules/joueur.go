package module

type Joueur struct {
	Nom   string
	Choix string // "X" ou "O"
}

// InitJoueur1 crée le joueur 1 avec son nom et son choix
func InitJoueur1(nom string, choix string) Joueur {
	return Joueur{
		Nom:   nom,
		Choix: choix,
	}
}

// InitJoueur2 crée le joueur 2 avec son nom et le symbole opposé au joueur 1
func InitJoueur2(nom string, choixJoueur1 string) Joueur {
	var choix string
	// Assigne le symbole opposé au joueur 1
	if choixJoueur1 == "X" {
		choix = "O"
	} else {
		choix = "X"
	}

	return Joueur{
		Nom:   nom,
		Choix: choix,
	}
}
