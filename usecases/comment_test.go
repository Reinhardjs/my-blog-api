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

func (s *Suite) Test_CommentUsecase_ReadById() {
	// Arrange
	s.repository.EXPECT().ReadById(1).Return(&models.Comment{
		ID:      1,
		Content: "This is content 1",
	}, nil)
	expectedComments := models.Comment{
		ID:      1,
		Content: "This is content 1",
	}

	// Act
	actualComments, err := s.usecase.ReadById(1)

	// Assert
	require.NoError(s.T(), err)
	require.Equal(s.T(), expectedComments, *actualComments)
}

func (s *Suite) Test_CommentUsecase_Create() {
	// Arrange
	newComment := &models.Comment{
		Content: "This-is-content",
	}
	s.repository.EXPECT().Create(newComment).Return(newComment, nil)

	// Act
	_, err := s.usecase.Create(newComment)

	// Assert
	require.NoError(s.T(), err)
}

func (s *Suite) Test_CommentUsecase_Update() {
	// Arrange
	const commentId = 1
	comment := &models.Comment{
		ID:      commentId,
		Content: "This-is-content",
	}
	s.repository.EXPECT().Update(commentId, comment).Return(comment, nil)

	// Act
	_, err := s.usecase.Update(commentId, comment)

	// Assert
	require.NoError(s.T(), err)
}
