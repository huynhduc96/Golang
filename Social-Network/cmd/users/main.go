package main

import (
	"log"
	"net/http"

	"social/internal/connector"
	"social/internal/constant"
	"social/internal/repository/cache"
	repository "social/internal/repository/database"
	"social/internal/repository/image"
	"social/internal/service"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Start Users Server \n")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Database connection
	db, err := connector.CreateDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cacheClient, err := connector.CreateRedisClient()
	if err != nil {
		log.Fatal(err)
	}
	defer cacheClient.Close()

	minIOClient, err := connector.CreateMinioClient()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("minIOClient: %+v\n", minIOClient)

	// Initialize repository
	userRepo := repository.CreateUserRepository(db)
	newsFeedRepo := repository.CreateNewsFeedRepository(db)
	sessionRepo := cache.CreateSessionRepository(cacheClient)
	imageStorage := image.CreateImageStorageRepository(minIOClient, constant.DefaultBucket)

	// Initialize service
	userService := service.NewUserService(*userRepo, *sessionRepo, *imageStorage)
	newsFeedService := service.NewNewsFeedService(*newsFeedRepo, *sessionRepo, *imageStorage)

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // Max 8MB

	// Health check
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	// User
	r.POST("/v1/user/login", userService.LoginUser) // login
	r.POST("/v1/user", userService.SignUpUsers)     // sign up
	r.PUT("/v1/user", userService.UpdateUser)       // update user info

	//  Profile
	r.GET("/v1/friends/:user_id", userService.GetFollowers)   // View Follow list
	r.POST("/v1/friends/:user_id", userService.FollowHandler) // Follow New User
	r.DELETE("/v1/friends/:id", userService.UnFollowHandler)  // Unfollow User

	// Posts
	r.GET("v1/friends/:user_id/posts", userService.GetUserPostByUserId) // View all User Posts by User Id
	r.POST("/v1/posts", userService.CreatePost)                         // Create Post
	r.POST("/v1/posts/:post_id/images", userService.UploadImage)        // create image
	r.PUT("/v1/posts/:post_id", userService.EditPost)                   // update post
	r.DELETE("/v1/posts/:post_id", userService.DeletePost)              // delete post
	r.GET("/v1/posts/:post_id", userService.GetPost)                    // get post by id

	// Comment
	r.POST("/v1/posts/:post_id/comments", userService.CommentPost) // create comment by post id
	r.GET("/v1/posts/:post_id/comments", userService.CommentPost)  // get all comment by post id

	// Like
	r.POST("/v1/posts/:post_id/likes", userService.LikePost)   // like/unlike by postId
	r.GET("/v1/posts/:post_id/likes", userService.GetLikePost) // get all like in post

	// NewFeeds
	r.GET("/v1/newsfeeds", newsFeedService.GetNewsFeed)

	r.Run()
}
