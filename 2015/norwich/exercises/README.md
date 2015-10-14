# Go Coding Exercises

To complete these excercises, you will need to look up documentation on the standard packages, which is available at "godoc.org/pkg".

However you should also be able to see these locally by running:
```
$ godoc -v -http=:6060 
```
Then you can see the documentation at "localhost:6060"


## Ex 1 - go-nude

go get github.com/koyachi/go-nude

Tasks: (A) What proportion of Ruben's pictures are judged nude? 
(B) How much faster does the task run when parallelized?


## Ex 2 - primesieve

Task: How fast does primesieve run with various numbers of cpus?

You may need the following functions in the standard packages:
```
runtime.NumCPU() int
runtime.GOMAXPROCS(n int)

time.Now() time.Time
time.Since(t time.Time) time.Duration
```

## Ex 3 - gowiki

Code is explained at https://golang.org/doc/articles/wiki/

To run it:
```
$ go run original.go
```
Then goto "http://localhost:8080/view/TestPage".

Task: Alter the code so that when a user goes to the root of the web server it does not give an error, 
instead list the wiki articles available and allow a new article to be added.

To start you off, here is an extract from a solution:
```
	http.HandleFunc("/", listTxtFiles)

	http.ListenAndServe(":8080", nil)
}

func listTxtFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>List of wiki entries</h1>")
```

You may need the following functions in the standard packages:
```
fd,err:=os.Open(".") // opens the current directory
fi,err:=fd.Readdir(-1) // reads all the entries in a directory

strings.HasSuffix()
strings.TrimSuffix()

os.Create()

(*http.Request).FormValue() 
http.Redirect()
```

## Ex 4 - gokit

The GoKit library is explained at gokit.io
```
$ go get github.com/go-kit/kit/examples/stringsvc1
$ stringsvc1
$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
{"v":"HELLO, WORLD","err":null}
$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
{"v":12}
```
Tasks: 
(A) Write a go client for stringvc1, creating uppercase() & count() functions work with any given string.
(B) Extend stringsvc1 to add a "lowercase" service, add use of the new service to your client.

You may find the following go functions helpful:
```
encoding/json - Marshal(), Unmarshal()
```

