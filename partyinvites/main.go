package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)

type Rsvp struct {
    Name, Email, Phone string
    WillAttend         bool
}

type formData struct {
    *Rsvp
    Errors []string
}

var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
    templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
    for index, name := range templateNames {
        t, err := template.ParseFiles("layout.html", name+".html")
        if err == nil {
            templates[name] = t

            fmt.Println("Loaded template", index, name)
        } else {
            panic(err)
        }
    }
}

func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
    err := templates["welcome"].Execute(writer, nil)
    if err != nil {
        http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
        log.Println("Template Execution Error", err)
        log.Printf("Error occurred while processing request: %v", request.URL)
        return
    }
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
    err := templates["list"].Execute(writer, responses)
    if err != nil {
        http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
        log.Println("Template Execution Error", err)
        log.Printf("Error occurred while processing request: %v", request.URL)
        return
    }
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
    if request.Method == http.MethodGet {
        err := templates["form"].Execute(writer, formData{
            Rsvp: &Rsvp{}, Errors: []string{},
        })
        if err != nil {
            http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
            log.Println("Template Execution Error", err)
            log.Printf("Error occurred while processing request: %v", request.URL)
            return
        }
    } else if request.Method == http.MethodPost {
        request.ParseForm()
        responseData := Rsvp{
            Name:       request.Form["name"][0],
            Email:      request.Form["email"][0],
            Phone:      request.Form["phone"][0],
            WillAttend: request.Form["willattend"][0] == "true",
        }

        errors := []string{}
        if responseData.Name == "" {
            errors = append(errors, "Please Enter Your Name")
        }
        if responseData.Email == "" {
            errors = append(errors, "Please enter your email address")
        }
        if responseData.Phone == "" {
            errors = append(errors, "Please enter your phone number")
        }
        if len(errors) > 0 {
            templates["form"].Execute(writer, formData{
                Rsvp: &responseData, Errors: errors,
            })
        } else {
            responses = append(responses, &responseData)
            if responseData.WillAttend {
                templates["thanks"].Execute(writer, responseData.Name)
            } else {
                templates["sorry"].Execute(writer, responseData.Name)
            }
        }
    }
}

func main() {
    loadTemplates()

    http.HandleFunc("/", welcomeHandler)
    http.HandleFunc("/list", listHandler)
    http.HandleFunc("/form", formHandler)

    err := http.ListenAndServe(":5000", nil)
    if err != nil {
        fmt.Println(err)
    }
}
