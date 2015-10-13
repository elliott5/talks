# Go Coding Exercises

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
fd.ReadDir(-1) // reads all the entries in a directory

os.Create()

strings.HasSuffix()
strings.TrimSuffix()

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
Task: write a go client for stringvc1

You may find the following go functions helpful:
```
encoding/json - Marshal(), Unmarshal()
net/http - Post()
bytes - NewReader()
io/ioutil - ReadAll()
```

