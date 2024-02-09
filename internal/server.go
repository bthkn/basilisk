package internal

import (
	"fmt"
	"net/http"

	"github.com/bthkn/basilisk/internal/appconfig"
)

type server struct {
	client *http.Client
	config *appconfig.AppConfig
}

type Server interface {
	Run() error
}

func (s server) Run() error {
	// return s.gin.Run(fmt.Sprintf("%s:%d", s.config.Host, s.config.Port))
	// return errors.New("oops")
	return http.ListenAndServe(":8080", nil)
}

func NewServer() Server { //config *appconfig.AppConfig
	s := &server{
		client: nil,
		// client: http.Server{Addr: "", Handler: "handler"},
	}
	return s
}

// from main.go :

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	// r.URL.Path
	// http.Redirect(w, r, "/edit/"+title, http.StatusFound)
}

// func serve(s *internal.Server) {
// var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
// m := validPath.FindStringSubmatch(r.URL.Path)
//   if m == nil {
//       http.NotFound(w, r)
//       return "", errors.New("invalid Page Title")
//   }
//   return m[2], nil // The title is the second subexpression.

// http.HandleFunc("/", handler)
// log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		m := validPath.FindStringSubmatch(r.URL.Path)
// 		if m == nil {
// 			http.NotFound(w, r)
// 			return
// 		}
// 		fn(w, r, m[2])
// 	}
// }

/*
cfg, err := appconfig.LoadFromPath(context.Background(), "pkl/dev/config.pkl")
if err != nil {
	panic(err)
}
if err = internal.NewServer(cfg).Run(); err != nil {
	panic(err)
}
*/
