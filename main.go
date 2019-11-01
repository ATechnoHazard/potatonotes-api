/*
 *     Copyright ATechnoHazard 2019  <amolele@gmail.com>.
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"github.com/ATechnoHazard/potatonotes-api/controllers"
	"github.com/ATechnoHazard/potatonotes-api/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.Use(middleware.JwtAuthentication) // use our jwt middleware

	// Auth routes
	router.HandleFunc("/api/users/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/users/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/users/delete", controllers.Delete).Methods("POST")
	router.HandleFunc("/api/users/info", controllers.UserInfo).Methods("GET")

	// User acc mgmt routes
	router.HandleFunc("/api/users/manage/username", controllers.ModifyUsername).Methods("POST")
	router.HandleFunc("/api/users/manage/password", controllers.ModifyPassword).Methods("POST")

	// Notes routes
	router.HandleFunc("/api/notes/create", controllers.CreateNote).Methods("POST")
	router.HandleFunc("/api/notes/list", controllers.ListNotes).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Printf("Listening on port %s\n", port)

	err := http.ListenAndServe(":"+port, router) // launch the middleware, visit localhost:4000/api
	if err != nil {
		log.Fatalln(err)
	}
}
