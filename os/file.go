package os

import "os"

func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}
}

func RemoveIfExists(path string) error {
	ok, err := FileExists(path)
	if err != nil {
		return err
	}
	if ok {
		return os.Remove(path)
	} else {
		return nil
	}
}
