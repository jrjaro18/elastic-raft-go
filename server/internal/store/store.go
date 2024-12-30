package store

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Store struct {
	data map[string]string
	logFile *os.File
}

func NewStore() *Store {
	err := os.Chdir("logs")
	if err != nil {
		log.Fatalln("Couldn't change the directory to logs\n", err)
	}
	file, err := os.OpenFile(uuid.New().String()+".log", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln("Couldn't open the store: reason unable to create file\n", err)
	}

	return &Store{
		data: make(map[string]string),
		logFile: file,
	}
}

func (s *Store) Contains(k string) bool {
	_, exists := s.data[k]
	return exists
}

func (s *Store) Add(k string, v string) error {
	s.data[k] = v
	_, err := fmt.Fprintln(s.logFile,"added<break>"+k+"<break>value<break>"+v)
	if err != nil {
		return errors.New("couldn't write in the file\n"+err.Error())
	}
	return nil
}

func (s *Store) Get(k string) (string, bool) {
	value, exists := s.data[k]
	return value, exists
}

func (s *Store) Remove(k string) error {
	delete(s.data, k)
	_, err := fmt.Fprintln(s.logFile,"removed<break>"+k)
	if err != nil {
		return errors.New("couldn't write in the file\n"+err.Error())
	}
	return nil
}

func (s *Store) Perform(opr string) error {
	if strings.HasPrefix(opr, "a") {
		x := strings.Split(opr, "<break>")
		return s.Add(x[1],x[2])
	} else {
		x := strings.Split(opr, "<break>")
		return s.Remove(x[1])
	}
}

func (s *Store) RebootLogFile() error {
	// clear the file
	err := os.Truncate(s.logFile.Name(), 0)
	if err != nil {
		return errors.New("couldn't clear the file\n"+err.Error())
	}
	s.logFile.Close()
	// open the file again, reason: reassigning the file pointer
	s.logFile, err = os.OpenFile(s.logFile.Name(), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return errors.New("couldn't open the file\n"+err.Error())
	}
	// add all the data to the file
	for k, v := range s.data {
		if err := s.Add(k, v); err != nil {
			return errors.New("couldn't add data to the file during reboot\n" + err.Error())
		}
	}
	return nil
}