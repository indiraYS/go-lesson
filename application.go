package main

import ("fmt"; "net/http"; "log"; "html/template")

type User struct {
    Name  string
    Age uint16
    Money int16
    Avg_grade, Happiness float64
    Hobbies []string
}

func (u User) getAllInfo() string {
  return fmt.Sprintf("Username is: %s. He is %d old and he has money equal: %d",
    u.Name, u.Age, u.Money)
}

func (u *User) setNewName(nName string) {
  u.Name = nName
}

func home_page(w http.ResponseWriter, r *http.Request) {

     bob := User{ "Bob", 25, -2,  4.2, 0.8, []string{}}
     bob.setNewName("Indira")

     //
     tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html")
     if err != nil {

         log.Println("yeah")
         fmt.Fprintf(w, err.Error())
         http.Error(w, "Internal Server Error", 500)
     }
     err =  tmpl.Execute(w, bob)
     if err != nil {
  	    log.Println(err.Error())
        http.Error(w, "Internal Server Error 3", 500)
     }
     http.Error(w, "Good", 500)

}

func contacts_page(w http.ResponseWriter, r *http.Request) {
     fmt.Fprintf(w, "Contacts page!")
}

func handleRequest() {
  http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir("./static/"))))
  http.HandleFunc("/", home_page)
  http.HandleFunc("/contacts/", contacts_page)
  http.ListenAndServe(":5000", nil)
}

// func dbConnect() {
//   psqlInfo := "postgres://2tors:2tors@localhost:5431/fork2tors"
//
//   db, err := pgx.Connect(context.Background(), psqlInfo)
//
//   if err != nil {
//     panic(err)
//   }
//
//   defer db.Close(context.Background())
//
//
//
//   err = conn.QueryRow(context.Background(), "select * where et_users", 42).Scan(&name, &weight)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 		os.Exit(1)
// 	}
//
// }

func main() {
     // var bob User = ...
     // bob := User{name: "Bob", age: 25, money: -2, avg_grade: 4.2, happiness: 0.8}
   handleRequest()
     // dbConnect()
     fmt.Println("Successfull connect")
}
