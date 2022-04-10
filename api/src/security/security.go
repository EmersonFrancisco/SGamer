package security

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma senha por string e coloca um Hash nela
func Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

// ValidadePass compara uma senha em String, e um hash e retorna se são iguais
func ValidadePass(passString, passHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passHash), []byte(passString))
}
