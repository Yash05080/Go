package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
	"github/gorilla/mux"

)
type Movie struct{
	ID string `json"ID"`
	Isbn string `"json"isbn"`
	Title string `json"title"`
	Director string `json"director"`
	
}

type Director struct{
	Firstname string `json"firstname"`
	Lastname string `json"lastname"`
}

var movies []Movie

func main(){
	r := mux.NewRouter("/movies")


}