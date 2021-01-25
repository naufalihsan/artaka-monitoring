package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gunturbudikurniawan/Artaka/api/auth"
	"github.com/gunturbudikurniawan/Artaka/api/models"
	"github.com/gunturbudikurniawan/Artaka/api/security"
	"github.com/gunturbudikurniawan/Artaka/api/utils/errors"
	"github.com/gunturbudikurniawan/Artaka/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) CreateMerchants(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	merchant := models.Subscribers{}

	err = json.Unmarshal(body, &merchant)
	if err != nil {
		restErr := errors.RestErr{
			Message: "Cannot unmarshal body",
			Status:  "Failed",
			Error:   "Unmarshal_error",
		}
		c.JSON(http.StatusBadRequest, restErr)
		return

	}
	merchant.Prepare()
	errorMessages := merchant.Validate("")
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	log.Println(err)
	userCreated, err := merchant.SaveUser(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"error":  formattedError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": userCreated,
	})
}

func (server *Server) LoginMerchant(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":      http.StatusUnprocessableEntity,
			"first error": "Unable to get request",
		})
		return
	}
	merchant := models.Subscribers{}
	err = json.Unmarshal(body, &merchant)
	log.Println(err)
	if err != nil {
		restErr := errors.RestErr{
			Message: "Cannot unmarshal body",
			Status:  "Failed",
			Error:   "Unmarshal_error",
		}
		c.JSON(http.StatusBadRequest, restErr)
		return

	}
	merchant.Prepare()
	errorMessages := merchant.Validate("login")
	if len(errorMessages) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errorMessages,
		})
		return
	}
	merchantData, err := server.SignInMerchant(merchant.Email, merchant.Secret_password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  formattedError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": merchantData,
	})
}

func (server *Server) SignInMerchant(email, password string) (map[string]interface{}, error) {

	var err error

	merchantData := make(map[string]interface{})

	merchant := models.Subscribers{}

	err = server.DB.Debug().Model(models.Subscribers{}).Where("email = ?", email).Take(&merchant).Error
	if err != nil {
		fmt.Println("this is the error getting the user: ", err)
		return nil, err
	}
	err = security.VerifyPassword(merchant.Secret_password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("this is the error hashing the password: ", err)
		return nil, err
	}
	// token, err := auth.CreateToken(merchant.ID)
	if err != nil {
		fmt.Println("this is the error creating the token: ", err)
		return nil, err
	}
	// merchantData["token"] = token
	merchantData["id"] = merchant.ID
	merchantData["email"] = merchant.Email

	return merchantData, nil
}

func (server *Server) UpdateMerchant(c *gin.Context) {

	errList = map[string]string{}

	userID := c.Param("id")
	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		restErr := errors.RestErr{
			Message: "Invalid Request",
			Status:  "Failed",
			Error:   "Invalid_request",
		}
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	tokenID, _, _, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"error":  errList,
		})
		return
	}
	if tokenID != 0 && tokenID != uint32(uid) {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"error":  errList,
		})
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusOK,
			"error":  errList,
		})
		return
	}

	requestBody := map[string]string{}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		restErr := errors.RestErr{
			Message: "Cannot unmarshal body",
			Status:  "Failed",
			Error:   "Unmarshal_error",
		}
		c.JSON(http.StatusBadRequest, restErr)
		return

	}
	formerMerchant := models.Subscribers{}
	err = server.DB.Debug().Model(models.Subscribers{}).Where("id = ?", uid).Take(&formerMerchant).Error
	if err != nil {
		errList["User_invalid"] = "The user is does not exist"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusOK,
			"error":  errList,
		})
		return
	}

	newMerchant := models.Subscribers{}

	if requestBody["current_password"] == "" && requestBody["new_password"] != "" {
		errList["Empty_current"] = "Please Provide current password"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusOK,
			"error":  errList,
		})
		return
	}
	if requestBody["current_password"] != "" && requestBody["new_password"] == "" {
		errList["Empty_new"] = "Please Provide new password"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusOK,
			"error":  errList,
		})
		return
	}
	if requestBody["current_password"] != "" && requestBody["new_password"] != "" {
		if len(requestBody["new_password"]) < 6 {
			errList["Invalid_password"] = "Password should be atleast 6 characters"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusOK,
				"error":  errList,
			})
			return
		}
		//if they do, check that the former password is correct
		err = security.VerifyPassword(formerMerchant.Secret_password, requestBody["current_password"])
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			errList["Password_mismatch"] = "The password not correct"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusOK,
				"error":  errList,
			})
			return
		}
		newMerchant.Owner_name = formerMerchant.Owner_name
		newMerchant.Email = requestBody["email"]
		newMerchant.Secret_password = requestBody["new_password"]
	}
	newMerchant.Owner_name = formerMerchant.Owner_name
	newMerchant.Email = requestBody["email"]

	newMerchant.Prepare()
	errorMessages := newMerchant.Validate("update")
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	updatedMerchant, err := newMerchant.UpdateMerchant(server.DB, uint32(uid))
	if err != nil {
		errList := formaterror.FormatError(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": updatedMerchant,
	})
}

func (server *Server) CreateSavedOrder(c *gin.Context) {

	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	savedorder := &models.Saved_orders{}
	err = json.Unmarshal(body, &savedorder)
	log.Println(err)
	if err != nil {
		restErr := errors.RestErr{
			Message: "Cannot unmarshal body",
			Status:  "Failed",
			Error:   "Unmarshal_error",
		}
		c.JSON(http.StatusBadRequest, restErr)
		return

	}
	uid, _, _, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		log.Println(err)
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	merchant := models.Subscribers{}
	err = server.DB.Debug().Model(models.Subscribers{}).Where("id = ?", uid).Take(&merchant).Error
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	savedorder.User_id = merchant.User_id
	savedorder.Prepare()

	errorMessages := savedorder.Validate()
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	orderCreated, err := savedorder.SaveOrder(server.DB)
	log.Println(err)
	if err != nil {
		log.Println(err)
		errList := formaterror.FormatError(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": orderCreated,
	})
}

func (server *Server) CreateSales(c *gin.Context) {

	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	saveSales := &models.Sales{}
	err = json.Unmarshal(body, &saveSales)
	if err != nil {
		restErr := errors.RestErr{
			Message: "Cannot unmarshal body",
			Status:  "Failed",
			Error:   "Unmarshal_error",
		}
		c.JSON(http.StatusBadRequest, restErr)
		return

	}
	uid, _, _, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	merchant := models.Subscribers{}
	err = server.DB.Debug().Model(models.Subscribers{}).Where("id = ?", uid).Take(&merchant).Error
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	saveSales.UserID = merchant.User_id
	saveSales.Prepare()

	errorMessages := saveSales.Validate()
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	salesCreated, err := saveSales.SaveSales(server.DB)
	if err != nil {
		errList := formaterror.FormatError(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": salesCreated,
	})
}

func (server *Server) CreateOnlineSales(c *gin.Context) {

	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	saveOnlineSales := &models.Onlinesales{}
	err = json.Unmarshal(body, &saveOnlineSales)
	log.Println(err)
	if err != nil {
		restErr := errors.RestErr{
			Message: "Cannot unmarshal body",
			Status:  "Failed",
			Error:   "Unmarshal_error",
		}
		c.JSON(http.StatusBadRequest, restErr)
		return

	}
	uid, _, _, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	merchant := models.Subscribers{}
	err = server.DB.Debug().Model(models.Subscribers{}).Where("id = ?", uid).Take(&merchant).Error
	if err != nil {

		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	saveOnlineSales.User_id = merchant.User_id
	saveOnlineSales.Prepare()

	errorMessages := saveOnlineSales.Validate()
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	salesOnlineCreated, err := saveOnlineSales.SaveOnlineSales(server.DB)
	if err != nil {
		log.Println(err)
		errList := formaterror.FormatError(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": salesOnlineCreated,
		"Error":    "Null",
	})
}
func (server *Server) GetCertainSubscribers(c *gin.Context) {
	// Is this user authenticated?
	_, referral_code, role, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	err, datas := models.ShowSubscribers(server.DB, referral_code, role)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "Failed",
			"error":    "Tidak ada merchants",
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

func (server *Server) GetMerchant(c *gin.Context) {

	errList = map[string]string{}

	userID := c.Param("id")

	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		restErr := errors.RestErr{
			Message: "Invalid Request",
			Status:  "Failed",
			Error:   "Invalid_request",
		}
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	merchant := models.Subscribers{}

	merchantGotten, err := merchant.FindMerchantByID(server.DB, uint32(uid))
	if err != nil {
		errList["No_user"] = "No User Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": merchantGotten,
	})
}

func (server *Server) GetLastMerchant(c *gin.Context) {

	transaction := models.Sales{}

	merchantLast, err := transaction.FindSales(server.DB)
	if err != nil {
		errList["No_user"] = "No User Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": merchantLast,
	})
}

func (server *Server) GetLastSaved(c *gin.Context) {

	transaction := models.Saved_orders{}

	merchantLast, err := transaction.FindSaved(server.DB)
	if err != nil {
		errList["No_user"] = "No User Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": merchantLast,
	})
}

func (server *Server) GetLastOnline(c *gin.Context) {

	transaction := models.Onlinesales{}

	merchantLast, err := transaction.FindOnline(server.DB)
	if err != nil {
		errList["No_user"] = "No User Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": merchantLast,
	})
}
