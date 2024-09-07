package service

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"social/internal/constant"
	"social/internal/models"
	"social/internal/repository/cache"
	repository "social/internal/repository/database"
	"social/internal/repository/image"
	"social/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserRepository performs database operations for users
type UserService struct {
	userRepository repository.UserRepository
	session        cache.SessionModel
	imageStorage   image.ImageStorageRepository
}

func NewUserService(userRepository repository.UserRepository, session cache.SessionModel, imageStorage image.ImageStorageRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
		session:        session,
		imageStorage:   imageStorage,
	}
}

func (r *UserService) getUserFromCookie(c *gin.Context) (*models.User, int, string) {
	cookie, err := c.Cookie("session")
	if err != nil {
		return nil, http.StatusUnauthorized, "Missing session cookie"
	}

	userID, err := r.session.GetUserIdFromSession(cookie)

	if err != nil || userID == -1 {
		return nil, http.StatusUnauthorized, "session cookie error"

	}

	curUser, err := r.userRepository.GetUserRecordByUserId(userID)

	if err != nil {
		return nil, http.StatusUnauthorized, "user not found"
	}

	return &curUser, http.StatusOK, ""
}

func (r *UserService) SignUpUsers(c *gin.Context) {
	var newUserRecord models.User
	err := c.ShouldBindBodyWithJSON(&newUserRecord)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	// Create new user
	bOBDate, err := utils.StringToTime(newUserRecord.DOB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	salt := utils.GenRandomSalt(constant.SaltSize)

	hashedPass := utils.HashPassword(newUserRecord.Password, salt)

	id, err := r.userRepository.CreateNewUser(&newUserRecord, bOBDate, salt, hashedPass)

	log.Printf("id: %+v\n", id)

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})

}

// LoginUser checks if the user exists in the database and returns it or an error
func (r *UserService) LoginUser(c *gin.Context) {
	var userLogin models.UserLogin
	err := c.ShouldBindBodyWithJSON(&userLogin)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	// get the user record from the database
	curUser, err := r.userRepository.GetUserRecordByUsername(userLogin.UserName)

	log.Printf("curUser: %+v\n", curUser)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusForbidden, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	if curUser.IsMatchPassword(userLogin.Password) == false {
		err = errors.New("wrong password")
		log.Println(err.Error())
		c.JSON(http.StatusForbidden, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	sessionID := "session-" + curUser.UserName

	err = r.session.CreateSession(sessionID, curUser.Id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	c.SetCookie("session", sessionID, constant.RedisTllMaxAge, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successfully",
	})
}

func (r *UserService) UpdateUser(c *gin.Context) {
	logInUser, statusCode, message := r.getUserFromCookie(c)

	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{"message": message})
		return
	}

	var updateUserRecord models.UserUpdate
	err := c.ShouldBindBodyWithJSON(&updateUserRecord)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	bOBDate, err := utils.StringToTime(updateUserRecord.DOB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	salt := utils.GenRandomSalt(constant.SaltSize)

	hashedPass := utils.HashPassword(updateUserRecord.Password, salt)

	err = r.userRepository.UpdateUser(logInUser.Id, &updateUserRecord, salt, hashedPass, bOBDate)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update successfully",
	})
}

func (r *UserService) GetFollowers(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	followers, err := r.userRepository.ViewFollowers(userId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"followers": followers,
	})
}

func (r *UserService) FollowHandler(c *gin.Context) {
	logInUser, statusCode, message := r.getUserFromCookie(c)

	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{"message": message})
		return
	}
	userId, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	err = r.userRepository.FollowUser(userId, logInUser.Id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Follow successfully",
	})
}

func (r *UserService) UnFollowHandler(c *gin.Context) {
	logInUser, statusCode, message := r.getUserFromCookie(c)

	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{"message": message})
		return
	}
	userId, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	err = r.userRepository.UnFollowHandler(userId, logInUser.Id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Follow successfully",
	})
}

func (r *UserService) GetUserPostByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	curUser, err := r.userRepository.GetUserRecordByUserId(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "user not found",
		})
		return
	}

	posts, err := r.userRepository.ViewFriendPost(curUser.Id)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	for _, post := range posts {
		imagePath, err := r.imageStorage.GetSignedUrl(post.ContentImagePath, constant.ImageURLMaxLive)
		if err != nil {
			log.Println(err.Error())
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

func (r *UserService) CreatePost(c *gin.Context) {
	logInUser, statusCode, message := r.getUserFromCookie(c)

	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{"message": message})
		return
	}

	var createPost models.CreatePost
	err := c.ShouldBindBodyWithJSON(&createPost)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	postId, err := r.userRepository.CreatePost(logInUser.Id, &createPost)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}
	// TODO
	// defer triggerGenNewsfeed(post.UserId)

	c.JSON(http.StatusOK, gin.H{
		"postId": postId,
	})
}

func (r *UserService) UploadImage(c *gin.Context) {
	_, statusCode, message := r.getUserFromCookie(c)
	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{"message": message})
		return
	}
	postId, err := strconv.Atoi(c.Param("post_id"))

	// single file
	file, header, err := c.Request.FormFile("filename")
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	filename := header.Filename

	log.Printf("filename: %+v\n", filename)
	//fSize := header.Size

	out, err := os.Create("./tmp/" + filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	defer out.Close()
	_, err = io.Copy(out, file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	fStat, err := out.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	// Move file cursor to the start, follow https://www.reddit.com/r/golang/comments/wv7hky/i_cant_upload_the_file_to_the_minio_bucket/
	_, err = out.Seek(0, io.SeekStart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	log.Printf("Stats %d\n", fStat.Size())
	imagePath, err := r.imageStorage.PutImage(out, postId, filename, fStat.Size())
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"imagePath": imagePath,
	})
}

func (r *UserService) EditPost(c *gin.Context) {
	logInUser, statusCode, message := r.getUserFromCookie(c)

	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{"message": message})
		return
	}

	var createPost models.CreatePost
	err := c.ShouldBindBodyWithJSON(&createPost)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(c.Param("post_id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	curPost, err := r.userRepository.GetPostById(postId)
	if err != nil || curPost.UserId != logInUser.Id {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	err = r.userRepository.UpdatePost(postId, &createPost)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	// TODO defer triggerGenNewsfeed(curPost.UserId)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (r *UserService) DeletePost(c *gin.Context) {
	logInUser, statusCode, message := r.getUserFromCookie(c)

	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{"message": message})
		return
	}

	postId, err := strconv.Atoi(c.Param("post_id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	curPost, err := r.userRepository.GetPostById(postId)
	if err != nil || curPost.UserId != logInUser.Id {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	err = r.userRepository.DeletePost(postId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (r *UserService) GetPost(c *gin.Context) {
	logInUser, statusCode, message := r.getUserFromCookie(c)

	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{"message": message})
		return
	}

	postId, err := strconv.Atoi(c.Param("post_id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	curPost, err := r.userRepository.GetPostById(postId)
	if err != nil || curPost.UserId != logInUser.Id {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	imagePath, err := r.imageStorage.GetSignedUrl(curPost.ContentImagePath, constant.ImageURLMaxLive)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	curPost.ContentImagePath = imagePath

	c.JSON(http.StatusOK, gin.H{
		"post": curPost,
	})
}

func (r *UserService) CommentPost(c *gin.Context) {
	logInUser, statusCode, message := r.getUserFromCookie(c)

	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{"message": message})
		return
	}

	postId, err := strconv.Atoi(c.Param("post_id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	curPost, err := r.userRepository.GetPostById(postId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	var cmt models.CreateComment
	err = c.ShouldBindBodyWithJSON(&cmt)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	cmtId, err := r.userRepository.CommentPost(curPost.Id, logInUser.Id, cmt.ContentText)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"commentId": cmtId,
	})
}

func (r *UserService) GetCommentPost(c *gin.Context) {
	_, statusCode, message := r.getUserFromCookie(c)

	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{"message": message})
		return
	}

	postId, err := strconv.Atoi(c.Param("post_id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	curPost, err := r.userRepository.GetPostById(postId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	comments, err := r.userRepository.ViewCommentPost(curPost.Id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})
}

func (r *UserService) LikePost(c *gin.Context) {
	logInUser, statusCode, message := r.getUserFromCookie(c)

	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{"message": message})
		return
	}

	postId, err := strconv.Atoi(c.Param("post_id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	curPost, err := r.userRepository.GetPostById(postId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	err = r.userRepository.LikePost(curPost.Id, logInUser.Id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (r *UserService) GetLikePost(c *gin.Context) {
	_, statusCode, message := r.getUserFromCookie(c)

	if statusCode != http.StatusOK {
		c.JSON(statusCode, gin.H{"message": message})
		return
	}

	postId, err := strconv.Atoi(c.Param("post_id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	curPost, err := r.userRepository.GetPostById(postId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	likes, err := r.userRepository.ViewLikePost(curPost.Id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"likes": likes,
	})
}
