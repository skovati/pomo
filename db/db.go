package db

import (
    "os"
    "encoding/json"
	"io/ioutil"

    "github.com/skovati/pomo/timer"
)

func GetTimer() (*timer.Timer, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return nil, err
    }

    ioutil.

    t := timer.NewTimer()

    return t, nil
    
}

func SaveTimer(t *timer.Timer) error {
    file, err := json.MarshalIndent(*t, "", " ")
    if err != nil {
        return err
    }
    path, err := os.UserHomeDir()
    err = ioutil.WriteFile(path + "/pomo.json", file, 0644)
    if err != nil {
        return err
    }
    return nil
}
