package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	// "time"
)

//==========================================================================\\
// var wg sync.WaitGroup
var mutex = &sync.RWMutex{}
var shalookup map[string]string
var md5lookup map[string]string

// func init() {
// 	// shalookup := make()
// }

func GuessSingle(sourceHash string, filename string) string {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var ret string = ""
	for scanner.Scan() {
		password := scanner.Text()

		// TODO - From the length of the hash you should know which one of these to check ...
		// add a check and logicial structure
		if len(sourceHash) == 32 {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
				ret = password
				return password
			}
		} else if len(sourceHash) == 64 {
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
				ret = password
				return password
			}
		} else {
			// fmt.Printf("[+] Password not found (MD5):")
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return ret
}

func GenHash(hashtype string, pass string) {
	// shalookup = make(map[string]string)
	// md5lookup = make(map[string]string)
	// wg.Add(2)
	// defer wg.Done()
	// log.Println("pass: ", pass)
	// time.Sleep(time.Second)
	if hashtype == "md5" {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(pass)))
		mutex.Lock()
		md5lookup[hash] = pass
		mutex.Unlock()

		// log.Println("in here\n\n\n")
	} else if hashtype == "sha256" {
		// time.Sleep(time.Second)
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(pass)))
		// mutex.Lock()
		shalookup[hash] = pass
		// mutex.Unlock()
	}
	// wg.Wait()
	// log.Println("sha is: ", shalookup)
}

func GenHashMaps(filename string) {

	shalookup = make(map[string]string)
	md5lookup = make(map[string]string)
	//TODO
	//itterate through a file (look in the guessSingle function above)
	//rather than check for equality add each hash:passwd entry to a map SHA and MD5 where the key = hash and the value = password
	//TODO at the very least use go subroutines to generate the sha and md5 hashes at the same time
	//OPTIONAL -- Can you use workers to make this even faster

	// var wg sync.WaitGroup

	// wg.Add(1)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		password := scanner.Text()

		GenHash("sha256", password)
		go GenHash("md5", password)

	}

	//TODO create a test in hscan_test.go so that you can time the performance of your implementation
	//Test and record the time it takes to scan to generate these Maps
	// 1. With and without using go subroutines
	// 2. Compute the time per password (hint the number of passwords for each file is listed on the site...)

}

func GetSHA(hash string) (string, error) {
	password, ok := shalookup[hash]
	if ok {
		log.Printf("[+] Password found (SHA256): %s\n", password)
		return password, nil

	} else {

		return "", errors.New("sha password does not exist")

	}
}

//TODO
func GetMD5(hash string) (string, error) {
	password, ok := md5lookup[hash]
	if ok {
		log.Printf("[+] Password found (MD5): %s\n", password)
		return password, nil

	} else {

		return "", errors.New("md5 password does not exist")

	}
}
