package controllers

import (
	"awesomeProject/api/dtos"
	"awesomeProject/api/response"
	"awesomeProject/dao"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Root(w http.ResponseWriter, r *http.Request) {
	response.SUCCESS(w, http.StatusOK, dtos.Result{Result: "FMB DUDE!"})
}

func (server *Server) AllPosts(w http.ResponseWriter, r *http.Request) {
	post := dao.Post{}
	//read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//define dto that we want to map with request body
	paginationDto := &dtos.LimitOffsetDto{}

	//map request body to dto object
	err = json.Unmarshal(body, paginationDto)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//receive fields
	offset := paginationDto.Offset
	limit := paginationDto.Limit
	//search in db
	posts, err := post.FindAllPostsWithLimit(server.DB, offset, limit)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.SUCCESS(w, http.StatusOK, posts)
}
func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	//read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//define dto that we want to map with request body
	post := dao.Post{}
	//map request body to dto object
	err = json.Unmarshal(body, &post)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//map with db table
	post.ToTable()
	//save to db
	postCreated, err := post.SavePost(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, postCreated.ID))
	response.SUCCESS(w, http.StatusCreated, postCreated)
}

func (server *Server) RatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	ratePost := dtos.RatePostDto{}

	err = json.Unmarshal(body, &ratePost)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post := dao.Post{}

	id := ratePost.Id
	rate := ratePost.Rate
	postUpdated, err := post.SetPostRate(server.DB, id, rate)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.SUCCESS(w, http.StatusCreated, postUpdated)

}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			log.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			log.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&dao.Post{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
