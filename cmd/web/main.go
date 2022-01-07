package pkg

import (
	"fmt"
	"log"
	"net/http"
	"github.com/NotAliAhmad/learningGO/pkg/handlers"
)

const portNumber = ":8080"


// main function 
func main(){

	//first arg is the path and second is the hander func
	http.HandleFunc("/",handlers.Homepage)
	http.HandleFunc("/about",handlers.Aboutpage)

	// control + c to to stop the service
	log.Println(fmt.Sprintf("Starting application on port %s",portNumber))
	// Starts the service
	_ = http.ListenAndServe(portNumber,nil) 
}
