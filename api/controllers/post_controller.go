package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gunturbudikurniawan/Artaka/api/models"

	"github.com/gin-gonic/gin"
	"github.com/gunturbudikurniawan/Artaka/api/auth"
	"github.com/gunturbudikurniawan/Artaka/api/utils/errors"
)

func (server *Server) CreatePost(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":   "Failed",
			"error":    "Failed",
			"Response": "Null",
		})
		return
	}
	post := new(models.Post)

	err = json.Unmarshal(body, &post)
	if err != nil {
		restErr := errors.RestErr{
			Message: "Cannot unmarshal body",
			Status:  "Failed",
			Error:   "Unmarshal_error",
		}
		c.JSON(http.StatusBadRequest, restErr)
		return

	}
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":   "Failed",
			"error":    "Failed",
			"Response": "Null",
		})
		return
	}

	user := models.Admin{}
	err = server.DB.Debug().Model(models.Admin{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		log.Println(err)
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":   "Failed",
			"error":    "Invalid Credentials",
			"Response": "Null",
		})
		return
	}

	post.AuthorID = uid //the authenticated user is the one creating the post

	post.Prepare()
	errorMessages := post.Validate()
	if len(errorMessages) > 0 {
		// errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":   "Failed",
			"error":    "Failed",
			"Response": "Null",
		})
		return
	}

	postCreated, err := post.SavePost(server.DB)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "Failed",
			"error":    "Invalid Credentials",
			"Response": "Null",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":   "Success",
		"response": postCreated,
	})
}

func (server *Server) GetPost(c *gin.Context) {

	postID := c.Param("id")
	pid, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	post := new(models.Post)

	postReceived, err := post.FindPostByID(server.DB, pid)
	if err != nil {
		errList["No_post"] = "No Post Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": postReceived,
	})
}

func (server *Server) UpdatePost(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	postID := c.Param("id")
	// Check if the post id is valid
	pid, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	// //CHeck if the auth token is valid and  get the user id from it
	// uid, err := auth.ExtractTokenID(c.Request)
	// if err != nil {
	// 	errList["Unauthorized"] = "Unauthorized"
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"status":   http.StatusUnauthorized,
	// 		"error":    "Unauthorized",
	// 		"response": "null",
	// 	})
	// 	return
	// }
	//Check if the post exist
	origPost := models.Post{}
	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&origPost).Error
	if err != nil {
		errList["No_post"] = "No Post Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status":   http.StatusNotFound,
			"error":    "No Post Found",
			"response": "null",
		})
		return
	}
	// if uid != origPost.AuthorID {
	// 	errList["Unauthorized"] = "Unauthorized"
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"status":   http.StatusUnauthorized,
	// 		"error":    "Unauthorized",
	// 		"Response": "null",
	// 	})
	// 	return
	// }
	// Read the data posted
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":   http.StatusUnprocessableEntity,
			"error":    "Unable to get request",
			"Response": "Null",
		})
		return
	}
	// Start processing the request data
	post := new(models.Post)
	err = json.Unmarshal(body, &post)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":   http.StatusUnprocessableEntity,
			"error":    "Cannot unmarshal body",
			"Response": "Null",
		})
		return
	}
	post.ID = origPost.ID //this is important to tell the model the post id to update, the other update field are set above
	post.AuthorID = origPost.AuthorID

	post.Prepare()
	errorMessages := post.Validate()
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":   http.StatusUnprocessableEntity,
			"error":    "Unauthorized",
			"Response": "Null",
		})
		return
	}
	postUpdated, err := post.UpdateAPost(server.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusInternalServerError,
			"error":    "Unauthorized",
			"response": "Null",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": postUpdated,
		"error":    "Null",
	})
}
func (server *Server) DeletePost(c *gin.Context) {

	postID := c.Param("id")
	// Is a valid post id given to us?
	pid, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	fmt.Println("this is delete post sir")

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	// Check if the post exist
	post := new(models.Post)
	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		errList["No_post"] = "No Post Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	// Is the authenticated user, the owner of this post?
	if uid != post.AuthorID {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	// If all the conditions are met, delete the post
	_, err = post.DeleteAPost(server.DB)
	if err != nil {
		errList["Other_error"] = "Please try again later"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": "Post deleted",
	})
}

func (server *Server) GetUserPosts(c *gin.Context) {

	userID := c.Param("id")
	// Is a valid user id given to us?
	uid, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	post := new(models.Post)
	posts, err := post.FindUserPosts(server.DB, uint32(uid))
	if err != nil {
		errList["No_post"] = "No Post Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": posts,
	})
}
func (server *Server) Showall(c *gin.Context) {

	err, datas := models.Allshow(server.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "Failed",
			"error":    "Merchant Aktif Semua",
			"response": "null",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "Success",
		"response": datas,
		"error":    "null",
	})
}
func (server *Server) ShowforWiranesia(c *gin.Context) {
	token := c.Request.Header.Get("authorization")
	fmt.Print("ini token", token)
	err, datas := models.ResponForWiranesia(server.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "Failed",
			"error":    "Tidak Ada Merchant Yang tidak respon",
			"response": "null",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "Success",
		"response": datas,
		"error":    "null",
	})
}

func (server *Server) LateRespon(c *gin.Context) {

	err, datas := models.NotRespon(server.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "Failed",
			"error":    "Tidak Ada Merchant Yang tidak respon",
			"response": "null",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "Success",
		"response": datas,
		"error":    "null",
	})
}

func (server *Server) LateResponWiranesia(c *gin.Context) {

	err, datas := models.NotResponWiranesia(server.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "Failed",
			"error":    "Tidak Ada Merchant Yang tidak respon",
			"response": "null",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "Success",
		"response": datas,
		"error":    "null",
	})
}

func (server *Server) NotAll(c *gin.Context) {

	err, datas := models.Show(server.DB)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "Failed",
			"error":    "Merchant Aktif Semua",
			"response": "null",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "Success",
		"response": datas,
		"error":    "null",
	})
}
func (server *Server) NotAllWiranesia(c *gin.Context) {

	err, datas := models.Showiranesia(server.DB)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "Failed",
			"error":    "Merchant Aktif Semua",
			"response": "null",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "Success",
		"response": datas,
		"error":    "null",
	})
}
func (server *Server) Already(c *gin.Context) {

	err, datas := models.Show1(server.DB)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "Failed",
			"error":    "Not ALready contacted with admin",
			"response": "null",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   "Success",
		"response": datas,
		"error":    "null",
	})
}

func (server *Server) AlreadyWiranesia(c *gin.Context) {

	err, datas := models.Show2(server.DB)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "Failed",
			"error":    "Not ALready contacted with admin",
			"response": "null",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   "Success",
		"response": datas,
		"error":    "null",
	})
}

func (server *Server) ShowSalesPayment(c *gin.Context) {

	err, datas := models.ShowPaymentMethodSales(server.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "Failed",
			"error":    "Tidak ada Payment Method Serupa",
			"response": "null",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "Success",
		"response": datas,
		"error":    "null",
	})
}

func (server *Server) ShowOnlineSalesPayment(c *gin.Context) {

	err, datas := models.ShowPaymentMethodVAOnlineSales(server.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "Failed",
			"error":    "Tidak ada Payment Method Serupa",
			"response": "null",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "Success",
		"response": datas,
		"error":    "null",
	})
}
