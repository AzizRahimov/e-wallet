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

//GetSha1 - генерация хеша
func GetSha1(text string, secret []byte) string {

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha1.New, secret)

	// Write Data to it
	h.Write([]byte(text))

	// Get result and encode as hexadecimal string
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}

// CalculateAmount - подсчет общей суммы с помощью горутин
func CalculateAmount(trn []models.Transaction) float64 {
	start := time.Now()

	var counter float64
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(len(trn))

	for _, value := range trn {

		value := value.Amount
		go func() {
			defer wg.Done()

			mu.Lock()

			counter += value
			mu.Unlock()

		}()

	}

	wg.Wait()

	fmt.Println(time.Now().Sub(start).Seconds())
	return counter
}
