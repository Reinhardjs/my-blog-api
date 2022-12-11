package repositories

import (
	"database/sql"
	"dot-crud-redis-go-api/models"
	"regexp"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type Suite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository CommentRepo
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)
	s.repository = CreateCommentRepo(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_CommentRepo_ReadAll() {
	// Arange
	var (
		id      = 1
		content = "this-is-content"
	)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "comments" WHERE "comments"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "content"}).
			AddRow(id, content))
	exp := []models.Comment{
		{
			ID:      1,
			Content: "this-is-content",
		},
	}

	// Act
	res, err := s.repository.ReadAll()

	// Assert
	require.NoError(s.T(), err)
	require.Equal(s.T(), exp, *res)
}

func (s *Suite) Test_CommentRepo_ReadById() {
	// Arrange
	const (
		id      = 1
		content = "this-is-content"
	)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "comments" WHERE "comments"."deleted_at" IS NULL`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "content"}).
			AddRow(id, content))
	exp := models.Comment{
		ID:      1,
		Content: "this-is-content",
	}

	// Act
	res, err := s.repository.ReadById(id)

	// Assert
	require.NoError(s.T(), err)
	require.Equal(s.T(), exp, *res)
}
