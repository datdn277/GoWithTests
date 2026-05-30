package dictionary

import "errors"

type Dictionary map[string]string

var (
	NotFoundError       = errors.New("could not find the word you were looking for")
	ErrWordExists       = errors.New("cannot add word because it already exists")
	ErrWordDoesNotExist = errors.New("cannot perform operation on word because it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]
	if !ok {
		return "", NotFoundError
	}
	return val, nil
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case NotFoundError:
		d[word] = def
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case NotFoundError:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}
