package main

import (
	"fmt"
	"github.com/srinathgs/mysqlstore"
	"net/http"
)

func main() {

	var store, _ = mysqlstore.NewMySQLStore(
		"root:root@tcp(localhost:3306)/mygoji?parseTime=true&loc=Local",
		"mysession",
		"/",
		3600,
		[]byte("12345678"))

	defer store.Close()

	var sessTest = func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "foobar")
		session.Values["bar"] = "baz"
		session.Values["baz"] = "foo"
		err = session.Save(r, w)

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%#v\n", session)
		fmt.Fprintln(w, err)

	}

	http.HandleFunc("/", sessTest)
	http.ListenAndServe(":8080", nil)
}
