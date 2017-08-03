// DGA Example project main.go
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	Word_A = [...]string{"try", "test", "bla"}
	Word_B = [...]string{"bla", "test", "try", ""}
	TLD    = [...]string{"com", "net", "org", "me", "tk"}

	PATH    string = "hi"
	Keyword string = "Hi"

	HASH bool = true
)

func md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	UsedArray := []string{}
	fmt.Println("Simple DGA (Domain Generation Algorithm)")
	fmt.Println("Possible Combinations:", len(Word_A)*len(Word_B)*len(TLD))
	var count int
	for i := 0; i < len(Word_A)*len(Word_B)*len(TLD); i++ {
	retry:
		var PartA string = Word_A[rand.Intn((len(Word_A)))]
		var PartB string = Word_B[rand.Intn((len(Word_B)))]
		var PartC string = TLD[rand.Intn((len(TLD)))]
		var Generated string
		if HASH {
			Generated = md5Hash(PartA+PartB) + "." + PartC
		} else {
			Generated = PartA + PartB + "." + PartC
		}
		for i := 0; i < len(UsedArray); i++ {
			if UsedArray[i] == Generated {
				goto retry
			}
		}
		UsedArray = append(UsedArray, Generated)
		count++
		resp, err := http.Get("http://" + Generated + "/" + PATH)
		if err != nil {
			goto retry
		}
		defer resp.Body.Close()
		html, err := ioutil.ReadAll(resp.Body)
		if err != nil {
\			goto retry
		}
		if string(html) == Keyword {
			fmt.Println("[Good]", "http://"+Generated, "["+string(strconv.Itoa(count))+"]")
		}
		if count >= len(Word_A)*len(Word_B)*len(TLD) {
			break
		}
	}
	fmt.Println("Job Done.")
}
