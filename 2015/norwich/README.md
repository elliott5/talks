# Norwich Go training day - 21st October 2015

Please follow the instructions below.

The first part of the day will speed through the [Go Tour](tour.golang.org), an introduction to the Go programming language, which is designed to be run online. 

However, as you will need Go installed for the second part of the day, you will be more self-sufficient if you install the Go Tour locally - here is how:

First [install Go 1.5.1](https://golang.org/doc/install), taking care to set the environment variable GOPATH correctly for your system. 

(On Windows, in the past I've needed to install github, and then use git-shell. But I'm not a Windows expert.)

Then install and run the Go Tour locally:
```
	$ go get golang.org/x/tour/gotour
	$ cd $GOPATH/bin
	$ ./gotour
```

Note that putting $GOPATH/bin in your PATH is helpful.

The second part of the day has some exercises, before the training day please type the following commands to get the required software onto your system (some of which will report "can't load package..." or "...no buildable Go source..."):

```
$ go get github.com/elliott5/talks/2015/norwich
$ go get github.com/koyachi/go-nude
$ go get github.com/go-kit/kit
```

Thank you, that will save time on the day.

If you have not watched it already, the [video on the front-page](https://vimeo.com/69237265) of [golang.org](golang.org) provides a good introduction to what we will be doing.

## Note: Please do not look at the solutions before you have attempted the exercises.

You will find the exercises at $GOPATH+"/src/github.com/elliott5/talks/2015/norwich/exercises"
