package db

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var bucketName = []byte("tasks")

// Task contains task data
type Task struct {
	ID   int
	Name string
}

// Init sets up the database
func Init() error {
	var err error
	db, err = bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}
		return nil
	})
}

// Tasks gets all tasks
func Tasks() []Task {
	var tasks []Task
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			id := btoi(k)
			tasks = append(tasks, Task{ID: id, Name: string(v)})
		}

		return nil
	})
	return tasks
}

// Add creates a new task
func Add(task string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		id, _ := b.NextSequence()
		key := itob(int(id))
		err := b.Put(key, []byte(task))
		if err != nil {
			return err
		}
		return nil
	})
}

// Delete removes a task by it's ID
func Delete(taskID int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		err := b.Delete(itob(taskID))
		if err != nil {
			return err
		}
		return nil
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi returns int representation of b
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
