package service

import (
	"log"
	"net/http"
	"social/internal/constant"
	"social/internal/models"
	"social/internal/repository/cache"
	repository "social/internal/repository/database"
	"social/internal/repository/image"

	"github.com/gin-gonic/gin"
)

type NewsFeedsService struct {
	newsFeedRepository repository.NewsFeedRepository
	session            cache.SessionModel
	imageStorage       image.ImageStorageRepository
}

func NewNewsFeedService(newsFeedRepository repository.NewsFeedRepository, session cache.SessionModel, imageStorage image.ImageStorageRepository) *NewsFeedsService {
	return &NewsFeedsService{
		newsFeedRepository: newsFeedRepository,
		session:            session,
		imageStorage:       imageStorage,
	}
}

func (r *NewsFeedsService) getUserFromCookie(c *gin.Context) (*models.User, int, string) {
	cookie, err := c.Cookie("session")
	if err != nil {
		return nil, http.StatusUnauthorized, "Missing session cookie"
	}

	userID, err := r.session.GetUserIdFromSession(cookie)

	if err != nil || userID == -1 {
		return nil, http.StatusUnauthorized, "session cookie error"

	}

	curUser, err := r.newsFeedRepository.GetUserRecordByUserId(userID)

	if err != nil {
		return nil, http.StatusUnauthorized, "user not found"
	}

	return &curUser, http.StatusOK, ""
}

func (r *NewsFeedsService) GetNewsFeed(c *gin.Context) {
	logInUser, statusCode, message := r.getUserFromCookie(c)

	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{"message": message})
		return
	}

	posts, err := r.newsFeedRepository.GenAllNewsFeedPost(logInUser.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}
	log.Printf("posts: %+v\n", posts)

	for _, post := range posts {
		log.Printf("ContentImagePath: %+v\n", post.ContentImagePath)
		imagePath, err := r.imageStorage.GetSignedUrl(post.ContentImagePath, constant.ImageURLMaxLive)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "err: " + err.Error(),
			})
			return
		}

		post.ContentImagePath = imagePath
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
