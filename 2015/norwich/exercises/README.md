# Go Coding Exercises

If you have some free time during the lunch break, try working through some example code at ["gobyexample.com"](https://gobyexample.com) or watch the short videos on the YouTube ["Go in 5 minutes"](https://www.youtube.com/channel/UC2GHqYE3fVJMncbrRd8AqcA/feed) channel. For more, there is a curated video selection at 


## Documentation

To complete some of the excercises, you will need to look up documentation on the standard packages, which is available at ["golang.org/pkg"](https://golang.org/pkg). However you should also be able to see these locally by running a Go documentation server:
```
$ godoc -http=:6060 &
```
Then you can see the documentation at "localhost:6060". It may take a few seconds to become available.


## Ex 1 - primesieve

The primeseive code is used as [an example](https://golang.org/ref/spec#An_example_package) in the Go specification, which is a really useful reference document. 

Tasks: 

(A) Limit the number of primes produced.

(B) How fast does the primesieve run with various numbers of CPUs? 

You may need the following functions in the standard packages:
```
runtime.NumCPU() int
runtime.GOMAXPROCS(n int)

time.Now() time.Time
time.Since(t time.Time) time.Duration
```


## Ex 2 - go-nude

Tasks: 

(A) What proportion of Ruben's pictures are judged nude? 

(B) How much faster does the task run when parallelized one goroutine per picture (unbounded parallelism)?

(C) How about one parallel stream per CPU (bounded parallelism)?

(D) How many parallel streams per CPU gives the same or better performance than one goroutine per picture?

You may find the "sync.WaitGroup" type helpful.


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

The GoKit library is explained at http://gokit.io

The server code is a slight adaptation of [stringsvc1](https://github.com/go-kit/kit/tree/master/examples).

```
$ go run *.go &
{"v":12}

$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
{"v":"HELLO, WORLD","err":null}

$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
{"v":12}
```

Tasks: 

(A) Write a go client for the built-in server, creating uppercase() & count() functions work with any given string.

You may find the following Go functions helpful:
```
encoding/json - Marshal(), Unmarshal()
```

(B) Extend stringsvc1 to add a "lowercase" service, add use of the new service to your client.

(C) Look at the stringsvc2 & stringsvc3 server code and extend them to use other go-kit facilities (not included in the provided solution).
