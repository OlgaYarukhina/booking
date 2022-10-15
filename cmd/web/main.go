package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/OlgaYarukhina/booking/pkg/config"
	"github.com/OlgaYarukhina/booking/pkg/handlers"
	"github.com/OlgaYarukhina/booking/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppCongig
var session *scs.SessionManager

//func handleRequest() {
//http.HandleFunc("/", handlers.Repo.Home)
//http.HandleFunc("/about", handlers.Repo.About)
//_ = http.ListenAndServe(potrNumber, nil)

func main() {

	//change this to true when in production
	app.InProguction = false

	// https://github.com/alexedwards/scs
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteDefaultMode
	session.Cookie.Secure = app.InProguction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create tamplate cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

	//handleRequest()
}

/*
//bob := User{"Bob", 25, -50, 4.2, 0.8, []string{}}
//tmpl, _ :=template.ParseFiles("templates/home_page.html")
	//tmpl.Execute(w,bob)
}

type User struct {
	Name string
	Age uint16 // невідємне
	Money int16
	Avg_grades, Happiness float64
	Hobbies []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money " +
		"equel: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string){
	u.Name = newName
	//bob.setNewName("Ivan")
	//fmt.Fprintf(w, bob.getAllInfo())
}

func addVal (x, y int) int{
	return x+y
}

//Сторінка роботи з помилками
http.HandleFunc("/divide", CheckErrors)

func CheckErrors(w http.ResponseWriter, r *http.Request) {
	f, err := divideVal(100.0, 10.0)
	if err != nil {
		fmt.Fprint(w, "Cannot divide by 0")
		return

	}
	fmt.Fprintf(w, fmt.Sprintf ("%f divide by %f is %f", 100.0, 10.0, f))
}

func divideVal (x,y float32) (float32, error){
	if y<=0 {
		err:= errors.New("Cannt")
		return 0, err
	}
	result:= x/y
	return result, nil
}


//Вивід тексту в браузер
func Home(w http.ResponseWriter, r *http.Request) {


	fmt.Fprintf(w, "This is home page \n")     //виводить на сторінку
	n, err := fmt.Fprintf(w, "Go is super easy 2") //виводить на сторінку
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Number of bytes: %d", n)) //виводить в термінал
	fmt.Print("Test") //виводить в термінал
}
*/
