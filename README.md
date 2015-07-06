## Double Array Trie

#### Intro

Double-Array Trie is a structure which is designed to make the size compact while maintaining fast access with algorithms for retrieval. It is more effective when using in Chinese/Japanese environment

Double-Array Trie is first presented in the paper below:

> [An Efficient Digital Search Algorithm by Using a Double-Array Structure](http://dl.acm.org/citation.cfm?id=75622)

reference: 
 [An Implementation of Double-Array Trie](http://linux.thai.net/~thep/datrie/datrie.html)

The project clone from the [go-darts](https://github.com/awsong/go-darts), but provides more feature

* remain the linked list trie when build double array trie
* generate output array for [go-darts](https://github.com/awsong/go-darts)

#### Usage

	package main

	import (
	    "bufio"
	    "bytes"
	    "fmt"
	    "io"
	    "os"
	    "gitlab.baidu.com/hanshinan/godarts"
	)

	func main() {
	    d := new(Darts)
	    dict, err := Read("test_keywords_eng")
	    if err != nil {
	        t.Error(err)
	    }

	    dat, _, err := d.Build(dict)
	    if err != nil {
		    fmt.Println(err)
	        return
	    }

	    //dat.PrintTrie() 	//double array trie
	    //llt.PrintTrie()	//linked list trie

	    for _, d := range dict {
	        if !dat.ExactMatchSearch(d, 0) {
		        fmt.Printf("%s not found\n", string(d))
	            return
	        }
	    }

	    fmt.Printf("Test total %d english words\n", len(dict))
	}
