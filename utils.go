package gommons

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func DeleteFile(path string) error {
    err := os.Remove(path)
    if err != nil {
        if os.IsNotExist(err) {
            err = nil
        }
    }
    return err
}

func IsSafeName(name string) bool {
    for _, v := range name {
        if (v < 'A' || v > 'Z') && (v < 'a' || v > 'z') && (v < '0' || v > '9') {
            switch v {
            case '_', '-':
                continue
            default:
                return false
            }
        }
    }

    return true
}

func CheckSafeName(name string) (err error) {
    if name == "" || !IsSafeName(name) {
        return fmt.Errorf("Invalid name: %s", name)
    }
    return nil
}

func ReadJson(path string, dest interface{}) error {
    bytes, err := ioutil.ReadFile(path)
    if err != nil {
        return err
    }

    err = json.Unmarshal(bytes, dest)
    if err != nil {
        log.Printf("Invalid JSON in file %s", path)
        return err
    }

    return nil
}
