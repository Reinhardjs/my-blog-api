package usecases

import (
	mocks "dot-crud-redis-go-api/mocks/repositories"
	"dot-crud-redis-go-api/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite

	ctrl       *gomock.Controller
	repository *mocks.MockCommentRepo
	usecase    CommentUsecase
}

func (s *Suite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.repository = mocks.NewMockCommentRepo(s.ctrl)
	s.usecase = CreateCommentUsecase(s.repository)
}

func (s *Suite) AfterTest(_, _ string) {
	// require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_CommentUsecase_ReadAll() {
	// Arrange
	s.repository.EXPECT().ReadAll().Return(&[]models.Comment{{
		ID:      1,
		Content: "This is content 1",
	}, {
		ID:      2,
		Content: "This is content 2",
	}}, nil)
	expectedComments := []models.Comment{
		{
			ID:      1,
			Content: "This is content 1",
		}, {
			ID:      2,
			Content: "This is content 2",
		},
	}

	// Act
	actualComments, err := s.usecase.ReadAll()

	// Assert
	require.NoError(s.T(), err)
	require.Equal(s.T(), expectedComments, *actualComments)
}
