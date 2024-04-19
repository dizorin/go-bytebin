package database

import (
	"github.com/dizorin/go-bytebin/database/cache"
	"github.com/dizorin/go-bytebin/models"
	"github.com/dizorin/go-bytebin/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

var (
	db = make(map[string][]byte)
)

// Save saves the data to the cache and db asynchronously
func Save(data []byte) (string, error) {
	key := utils.GenerateID()
	cf := utils.NewCompletableFuture[*models.Content]()

	cache.Set(key, cf)
	//_, err := executor.Scheduler.Do(saveContent, cf, key, data)
	go func() {
		saveContent(cf, key, data)
	}()
	//if err != nil {
	//	log.Warnf("Save data[key=%s]: Run executor %v", key, err)
	//}

	return key, nil
}

// Load retrieves content from cache and db if not found in cache
func Load(key string) (*models.Content, bool) {
	cf, found := cache.Get(key)
	if !found {
		cf = utils.NewCompletableFuture[*models.Content]()
		go func() {
			loadContent(cf, key)
		}()
	} else {
		cache.Update(key, cf)
	}

	get, err := cf.CompleteGetWithTimeout(utils.GetenvDuration("WAIT_LOAD_CONTENT"))

	if err != nil {
		return nil, false
	}
	return get.(*models.Content), true
}

// loadContent loads content from the db and saves it to the cache
func loadContent(cf *utils.CompletableFuture[*models.Content], key string) {
	time.Sleep(time.Second * 5) // TODO remove

	data, found := db[key]
	if !found {
		log.Warnf("Load data[key=%s]: not found", key)
		cf.Complete(nil, fiber.ErrNotFound)
	} else {
		content := &models.Content{Body: data}
		cf.Complete(content, nil)

		log.Infof("Loaded data[key=%s]", key)
		cache.Set(key, cf)
	}
}

// saveContent saves the content to the db and saves it to the cache
func saveContent(cf *utils.CompletableFuture[*models.Content], key string, data []byte) {
	time.Sleep(time.Second * 5) // TODO remove

	db[key] = data
	content := &models.Content{Body: data}

	cf.Complete(content, nil)

	cache.Set(key, cf)
	log.Infof("Saved data[key=%s]", key)
}
