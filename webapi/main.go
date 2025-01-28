package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type whoami struct {
    Name  string
    Title string
    State string
}

func main() {
    request1()
}

func whoAmI(response http.ResponseWriter, r *http.Request) {
    who := []whoami{
        {
            Name:  "Les bests du devops et les 2 autres",
            Title: "DevOps and Continuous Deployment",
            State: "FR",
        },
    }

    json.NewEncoder(response).Encode(who)
    fmt.Println("Endpoint Hit: /whoami")
}

func homePage(response http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(response, "Welcome to the Web API!")
    fmt.Println("Endpoint Hit: /")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
    who := "EfreiParis"
    fmt.Fprintf(response, "A little bit about me...")
    fmt.Println("Endpoint Hit: /aboutme ->", who)
}

func errorHandler(response http.ResponseWriter, r *http.Request) {
    http.Error(response, "Une erreur est survenue", http.StatusInternalServerError)
    fmt.Println("Endpoint Hit: /error (erreur générée)")
}

func request1() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/aboutme", aboutMe)
    http.HandleFunc("/whoami", whoAmI)
    
    http.HandleFunc("/error", errorHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
