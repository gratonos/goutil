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
	err := os.Remove(path)
	if os.IsNotExist(err) {
		return nil
	} else {
		return err
	}
}
