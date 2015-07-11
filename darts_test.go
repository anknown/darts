package godarts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func Read(filename string) ([][]rune, error) {
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

func TestBuild(t *testing.T) {
	d := new(Darts)
	dict, err := Read("test_keywords_eng")
	if err != nil {
		t.Error(err)
		return
	}

	_, _, err = d.Build(dict)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestBuildWithNothing(t *testing.T) {
	d := new(Darts)
	_, _, err := d.Build(nil)
	if err == nil {
		t.Error("test build without nothing failed")
		return
	}
}

func TestSearchEnglish(t *testing.T) {
	d := new(Darts)
	dict, err := Read("test_keywords_eng")
	if err != nil {
		t.Error(err)
		return
	}

	dat, _, err := d.Build(dict)
	if err != nil {
		t.Error(err)
		return
	}

	//dat.PrintTrie()
	//llt.PrintTrie()

	for _, d := range dict {
		if !dat.ExactMatchSearch(d, 0) {
			t.Errorf("%s not found", string(d))
			return
		}
	}

	fmt.Printf("Test total %d english words\n", len(dict))
}

func TestSearchChinese(t *testing.T) {
	d := new(Darts)
	dict, err := Read("test_keywords_chn")
	if err != nil {
		t.Error(err)
		return
	}

	dat, _, err := d.Build(dict)
	if err != nil {
		t.Error(err)
		return
	}

	//dat.PrintTrie()
	//llt.PrintTrie()

	for _, d := range dict {
		if !dat.ExactMatchSearch(d, 0) {
			t.Errorf("%s not found", string(d))
			return
		}
	}

	fmt.Printf("Test total %d chinese words\n", len(dict))
}
