package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"task/helper"
	"task/middleware"
	"task/model/domain"
	"task/model/web"
	"task/service"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaims struct {
	Id       int    `json:id`
	Name     string `json:"name"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthUserHandler interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindCareer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindCareerById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type UserController struct {
	authService service.AuthService
}

func NewAuthController(service service.AuthService) UserController {
	return UserController{
		authService: service,
	}
}

func (a *UserController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &createRequest)

	response := a.authService.Register(request.Context(), createRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   response,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (a *UserController) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	createRequest := web.LoginRequest{}
	helper.ReadFromRequestBody(request, &createRequest)

	hashPass, err := a.authService.Login(request.Context(), createRequest)
	helper.PanicHandling(err)

	if err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(createRequest.Password)); err != nil {
		webResponse := web.WebResponse{
			Code:   500,
			Status: "Failed to Login",
			Data:   "User or Password doesnt match",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	userDomain := a.authService.CheckUsername(request.Context(), createRequest.Username)

	claims := &jwtCustomClaims{
		userDomain.Id,
		userDomain.Name,
		userDomain.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	helper.PanicHandling(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Found",
		Data:   t,
	}
	helper.WriteToResponseBody(writer, webResponse)

}

func (a *UserController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (a *UserController) FindCareerById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authToken := request.Header.Get("Authorization")
	_, isTrue := middleware.ExtractClaims(authToken)
	if authToken == "" || isTrue == false {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   "Invalidate Token",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	res, err := http.Get("http://dev3.dansmultipro.co.id/api/recruitment/positions.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)

	var careers []domain.Career
	json.Unmarshal(bodyBytes, &careers)

	careerId := params.ByName("id")
	helper.PanicHandling(err)
	for _, value := range careers {
		if value.Id == careerId {
			webResponse := web.WebResponse{
				Code:   200,
				Status: "success",
				Data:   value,
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
	}

	webResponse := web.WebResponse{
		Code:   404,
		Status: "not found",
		Data:   nil,
	}
	helper.WriteToResponseBody(writer, webResponse)

}

func (a *UserController) FindCareer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authToken := request.Header.Get("Authorization")
	_, isTrue := middleware.ExtractClaims(authToken)
	if authToken == "" || isTrue == false {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   "Invalidate Token",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	res, err := http.Get("http://dev3.dansmultipro.co.id/api/recruitment/positions.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)

	var careers []domain.Career
	json.Unmarshal(bodyBytes, &careers)

	location := request.URL.Query().Get("location")
	full_time := request.URL.Query().Get("full_time")
	description := request.URL.Query().Get("description")
	page := request.URL.Query().Get("page")

	position := "full time"
	if full_time != "true" {
		position = "part time"
	}

	if location == "" && full_time == "" && description == "" {
		if page != "" {
			page, _ := strconv.Atoi(page)
			start := page*10 - 10
			end := page*10 - 1
			if len(careers) < end {
				end = len(careers)
			}
			webResponse := web.WebResponse{
				Code:   200,
				Status: "success",
				Data:   careers[start:end],
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
		webResponse := web.WebResponse{
			Code:   200,
			Status: "success",
			Data:   careers,
		}
		helper.WriteToResponseBody(writer, webResponse)

	}

	var filtering []domain.Career
	for _, value := range careers {
		if strings.ToLower(value.Location) == strings.ToLower(location) {
			filtering = append(filtering, value)
			continue
		}

		isMatch, _ := regexp.MatchString("\\b"+strings.ToLower(description)+"\\b", strings.ToLower(value.Description))

		if description != "" && isMatch {
			filtering = append(filtering, value)
			continue
		}
		if strings.ToLower(value.Type) == position {
			filtering = append(filtering, value)
			continue
		}
	}
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   filtering,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
