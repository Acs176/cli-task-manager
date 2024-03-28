package bolt

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"

	"github.com/Acs176/cli-task-manager/tasks"
	"github.com/boltdb/bolt"
)

type TaskDB struct {
	db *bolt.DB
}

func New() *TaskDB {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucket([]byte("Tasks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return &TaskDB{db}
}

func (tdb *TaskDB) AddTask(taskName string) error {
	return tdb.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		id, _ := b.NextSequence()
		err := b.Put(itob(id), []byte(taskName))
		return err
	})
}

func (tdb *TaskDB) DeleteTask(id int) error {
	return tdb.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		t := b.Get(itob(uint64(id)))
		if t == nil {
			return errors.New("the task does not exist")
		}
		err := b.Delete(itob(uint64(id)))
		return err
	})
}

func (tdb *TaskDB) ListTasks() []*tasks.Task {
	taskSlice := make([]*tasks.Task, 0)
	tdb.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			t := tasks.Task{Name: string(v), Index: binary.BigEndian.Uint64(k)}
			taskSlice = append(taskSlice, &t)
		}
		return nil
	})
	return taskSlice
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func btoi(v []byte) uint64 {
	s := string(v)
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return uint64(i)
}
