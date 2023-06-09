package persistence

import (
	"github.com/jesusEstaba/calculator/internal/database"
	"github.com/jesusEstaba/calculator/pkg/domain"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type UserRepositoryImplementation struct {
	repository database.CRUDRepository[domain.User, *domain.User]
}

func NewUserRepository(db *mongo.Database) domain.UserRepository {
	repository, err := database.NewRepository[domain.User](db, "users")
	if err != nil {
		logrus.Fatal(err)
	}

	return &UserRepositoryImplementation{
		repository: repository,
	}
}

func (r *UserRepositoryImplementation) GetUser(userID string) (*domain.User, error) {
	user, err := r.repository.FindByID(userID)

	return user, err
}

func (r *UserRepositoryImplementation) FindByUsername(username string) (*domain.User, error) {
	user, err := r.repository.FindByKey("username", username)
	if err != nil && strings.Contains(err.Error(), "no documents in result") {
		return nil, nil
	}

	return user, err
}

func (r *UserRepositoryImplementation) Save(user domain.User) (domain.User, error) {
	err := r.repository.Insert(&user)

	return user, err
}

func (r *UserRepositoryImplementation) UpdateUser(user domain.User) error {
	err := r.repository.Insert(&user)

	return err
}
