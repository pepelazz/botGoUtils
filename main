package goBotUtils

import "os"

func CreateFile(path string) (err error) {
	_, err = os.Stat(path)

	if os.IsNotExist(err) {
		err = nil
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	} else {
		DeleteFile(path)
		return CreateFile(path)
	}
	return
}

func DeleteFile(path string) error {
	return os.Remove(path)
}
