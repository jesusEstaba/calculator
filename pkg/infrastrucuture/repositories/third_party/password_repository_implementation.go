package third_party

import (
	"errors"
	"github.com/jesusEstaba/calculator/pkg/domain"
	"golang.org/x/crypto/bcrypt"
)

type PasswordRepositoryImplementation struct{}

func NewPasswordRepository() domain.PasswordRepository {
	return &PasswordRepositoryImplementation{}
}

func (r *PasswordRepositoryImplementation) Generate(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (r *PasswordRepositoryImplementation) Compare(passwd string, storedPasswd []byte) error {
	err := bcrypt.CompareHashAndPassword(storedPasswd, []byte(passwd))
	if err != nil {
		return errors.New("passwords does not match")
	}

	return nil
}
