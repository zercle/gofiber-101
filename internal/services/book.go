package services

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zercle/gofiber-101/internal/datasources"
	"github.com/zercle/gofiber-101/pkg/models"
	"gorm.io/gorm"
)

// fiber's router will parse fiber.Ctx pointer into handler
// that we can use for get request, send response or stores variables to next routes
// https://docs.gofiber.io/api/ctx
func GetBooks(ctx *fiber.Ctx) (err error) {
	// response model for response to client
	var responseForm models.ResponseForm
	// init err response variable incase that need to response what's wrong to client
	var errRespArr []models.ResposeError
	// init array of Book to query and store value
	var books []models.Book
	// use database connection from datasources.DBConn
	db := datasources.DBConn
	// gorm will query all books by Book model
	// https://gorm.io/docs/query.html#Retrieving-all-objects
	err = db.Find(&books).Error
	// incase want to debug query statement just add Debug()
	// https://gorm.io/docs/logger.html#Debug
	// err = db.Debug().Find(&books).Error
	// if some error set HTTP status code and return
	if err != nil {
		// log message on server console
		log.Printf("GetBooks err: %+v", err)
		// set HTTP status code
		// https://docs.gofiber.io/api/ctx#status
		ctx.Status(http.StatusInternalServerError)
		// fill error response body
		errRespObj := models.ResposeError{
			// http constant from go http package
			// https://pkg.go.dev/net/http?utm_source=gopls#pkg-constants
			Code:   http.StatusInternalServerError,
			Source: "GetBooks",
			// convert http status code to text
			// https://pkg.go.dev/net/http?utm_source=gopls#StatusText
			Title: http.StatusText(http.StatusInternalServerError),
			// get error message from err object
			Message: err.Error(),
		}
		// append error to error respons array
		errRespArr = append(errRespArr, errRespObj)
		// response error data in json format
		// datatype interface{} that use when we can't fix data type
		return ctx.JSON(map[string]interface{}{"errors": errRespArr})
	}

	responseForm = models.ResponseForm{
		Success: bool(err == nil),
		// we can minimize map[string]interface{} to fiber.Map{}
		Result: fiber.Map{"books": books},
	}
	// response book data in json format
	return ctx.JSON(responseForm)
}

func GetBook(ctx *fiber.Ctx) (err error) {
	var responseForm models.ResponseForm
	var errRespArr []models.ResposeError

	// get path param from fiber's context
	id := ctx.Params("id")
	// due to ctx.Params always return string
	// so we must check & convert into need datatype
	intID, err := strconv.Atoi(id)
	// if user fill wrong data type just warn
	if err != nil {
		log.Printf("GetBook err: %+v", err)
		ctx.Status(http.StatusBadRequest)
		errRespObj := models.ResposeError{
			Code:    http.StatusBadRequest,
			Source:  "GetBook",
			Title:   http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		}
		errRespArr = append(errRespArr, errRespObj)
		return ctx.JSON(fiber.Map{"errors": errRespArr})
	}

	var book models.Book
	db := datasources.DBConn
	// query with condition
	// https://gorm.io/docs/query.html#Conditions
	err = db.Where(models.Book{ID: uint(intID)}).Find(&book).Error
	if err != nil {
		log.Printf("GetBook err: %+v", err)
		ctx.Status(http.StatusInternalServerError)
		errRespObj := models.ResposeError{
			Code:    http.StatusInternalServerError,
			Source:  "GetBook",
			Title:   http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		}
		errRespArr = append(errRespArr, errRespObj)
		return ctx.JSON(fiber.Map{"errors": errRespArr})
	}

	responseForm = models.ResponseForm{
		Success: bool(err == nil),
		Result:  fiber.Map{"book": book},
	}
	return ctx.JSON(responseForm)
}

func NewBook(ctx *fiber.Ctx) (err error) {
	var responseForm models.ResponseForm
	var errRespArr []models.ResposeError
	book := new(models.Book)
	// parse request body into book model
	if err := ctx.BodyParser(book); err != nil {
		log.Printf("NewBook err: %+v", err)
		ctx.Status(http.StatusUnprocessableEntity)
		errRespObj := models.ResposeError{
			Code:    http.StatusUnprocessableEntity,
			Source:  "NewBook",
			Title:   http.StatusText(http.StatusUnprocessableEntity),
			Message: err.Error(),
		}
		errRespArr = append(errRespArr, errRespObj)
		return ctx.JSON(fiber.Map{"errors": errRespArr})
	}

	db := datasources.DBConn
	// use database transaction that we can rollback when something go wrong else commit when success
	// https://gorm.io/docs/transactions.html#Control-the-transaction-manually
	dbTx := db.Begin()
	// go defer will call by the end of function scope
	// https://gobyexample.com/defer
	defer dbTx.Rollback()

	// creat a record data
	// https://gorm.io/docs/create.html
	if err := dbTx.Create(&book).Error; err != nil {
		log.Printf("NewBook err: %+v", err)
		ctx.Status(http.StatusInternalServerError)
		errRespObj := models.ResposeError{
			Code:    http.StatusInternalServerError,
			Source:  "NewBook",
			Title:   http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		}
		errRespArr = append(errRespArr, errRespObj)
		return ctx.JSON(fiber.Map{"errors": errRespArr})
	}

	// commit & end the transaction
	err = dbTx.Commit().Error
	if err != nil {
		log.Printf("NewBook err: %+v", err)
		ctx.Status(http.StatusInternalServerError)
		errRespObj := models.ResposeError{
			Code:    http.StatusInternalServerError,
			Source:  "NewBook",
			Title:   http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		}
		errRespArr = append(errRespArr, errRespObj)
		return ctx.JSON(fiber.Map{"errors": errRespArr})
	}

	responseForm = models.ResponseForm{
		Success: bool(err == nil),
		Result:  fiber.Map{"book": book},
	}
	return ctx.JSON(responseForm)
}

func DeleteBook(ctx *fiber.Ctx) (err error) {
	var responseForm models.ResponseForm
	var errRespArr []models.ResposeError

	id := ctx.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("DeleteBook err: %+v", err)
		ctx.Status(http.StatusBadRequest)
		errRespObj := models.ResposeError{
			Code:    http.StatusBadRequest,
			Source:  "DeleteBook",
			Title:   http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		}
		errRespArr = append(errRespArr, errRespObj)
		return ctx.JSON(fiber.Map{"errors": errRespArr})
	}

	var book models.Book
	db := datasources.DBConn

	dbTx := db.Begin()
	defer dbTx.Rollback()

	err = dbTx.Where(models.Book{ID: uint(intID)}).First(&book).Error
	// check is record not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(404).SendString("")
	}
	// real error from database
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("DeleteBook err: %+v", err)
		ctx.Status(http.StatusInternalServerError)
		errRespObj := models.ResposeError{
			Code:    http.StatusInternalServerError,
			Source:  "DeleteBook",
			Title:   http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		}
		errRespArr = append(errRespArr, errRespObj)
		return ctx.JSON(fiber.Map{"errors": errRespArr})
	}

	err = dbTx.Where(models.Book{ID: uint(intID)}).Delete(&book).Error
	if err != nil {
		log.Printf("DeleteBook err: %+v", err)
		ctx.Status(http.StatusInternalServerError)
		errRespObj := models.ResposeError{
			Code:    http.StatusInternalServerError,
			Source:  "DeleteBook",
			Title:   http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		}
		errRespArr = append(errRespArr, errRespObj)
		return ctx.JSON(fiber.Map{"errors": errRespArr})
	}

	err = dbTx.Commit().Error
	if err != nil {
		log.Printf("DeleteBook err: %+v", err)
		ctx.Status(http.StatusInternalServerError)
		errRespObj := models.ResposeError{
			Code:    http.StatusInternalServerError,
			Source:  "DeleteBook",
			Title:   http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		}
		errRespArr = append(errRespArr, errRespObj)
		return ctx.JSON(fiber.Map{"errors": errRespArr})
	}

	responseForm = models.ResponseForm{
		Success: bool(err == nil),
		Result:  fiber.Map{"book": book},
	}
	return ctx.JSON(responseForm)
}
