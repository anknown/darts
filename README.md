## Double Array Trie

#### Intro

Double-Array Trie is a structure designed to make the size compact while maintaining fast access with algorithms for retrieval. It is more effective when using in Chinese/Japanese environment

Double-Array Trie is first presented in the paper below:

> [An Efficient Digital Search Algorithm by Using a Double-Array Structure](http://dl.acm.org/citation.cfm?id=75622)

reference: 
 [An Implementation of Double-Array Trie](http://linux.thai.net/~thep/datrie/datrie.html)

The project clones from the [go-darts](https://github.com/awsong/go-darts), but provides two more features

* remain the linked list trie when build double array trie
* generate output array for [aho-corasick algorithm](https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm)

#### Usage

	package main
	
	import (
	    "bufio"
	    "bytes"
	    "fmt"
	    "io"
	    "os"
	)
	
	import (
	    "github.com/anknown/darts"
	)
	
	func ReadRunes(filename string) ([][]rune, error) {
	    dict := [][]rune{}
	
	    f, err := os.OpenFile(filename, os.O_RDONLY, 0660)
	    if err != nil {
	        return nil, err
	    }
	
	    r := bufio.NewReader(f)
	    for {
	        l, err := r.ReadBytes('\n')
	        if err != nil || err == io.EOF {
	            break
	        }
	        l = bytes.TrimSpace(l)
	        dict = append(dict, bytes.Runes(l))
	    }
	
	    return dict, nil
	}

	func main() {
	    d := new(godarts.Darts)
	    dict, err := ReadRunes("your dict file")
	    if err != nil {
	        fmt.Println(err)
	        return
	    }
	
	    dat, _, err := d.Build(dict)
	    if err != nil {
	        fmt.Println(err)
	        return
	    }
	
	    //dat.PrintTrie()   //double array trie
	    //llt.PrintTrie()   //linked list trie
	
	    for _, d := range dict {
	        if !dat.ExactMatchSearch(d, 0) {
	            fmt.Printf("%s not found\n", string(d))
	            return
	        }
	    }
	
	    fmt.Printf("Test total %d english words\n", len(dict))
	}

#### License

MIT License
