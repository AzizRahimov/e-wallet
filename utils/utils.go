package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/AzizRahimov/e-wallet/models"
	"sync"
	"time"
)

func GetSha1(text string, secret []byte) string {

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha1.New, secret)

	// Write Data to it
	h.Write([]byte(text))

	// Get result and encode as hexadecimal string
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}

func WriteWithMutex(trn []models.Transaction) float64 {
	start := time.Now()
	//total := []int{10, 10, 10, 20, 50, 20, 50, 30}
	var counter float64
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(len(trn))

	//for i := 0; i < 10; i++ {
	for _, value := range trn {

		value := value.Amount
		go func() {
			defer wg.Done()

			mu.Lock()
			//! в данном участке кода мы может быть уверены, что работает только 1 горутина
			//! далее после того как мы внесли какие-то изменения мы разблокируем
			counter += value
			mu.Unlock()

		}()

	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
	return counter
}
