package third_party

import (
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
