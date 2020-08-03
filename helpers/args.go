package helpers
 
import "flag"

type Args struct {
	Message string		`json:"firstname,omitempty"`
	Path  string 		`json:"lastname,omitempty"`
}


func Create()  *Args{

	// https://golang.org/pkg/flag/
	// https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go-pt
	// https://gobyexample.com/command-line-flags
	flag.Parse()

	return '', ''
}

// voisynt --message "meutexto" --path "C:/www/enbot/encore/audio" --