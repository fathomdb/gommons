package gommons

import (
    "io/ioutil"
    "log"
    "os"
    "sort"
    "syscall"
)

func StatIfExists(path string) (stat os.FileInfo, err error) {
    stat, err = os.Stat(path)
    if err != nil {
        if os.IsNotExist(err) {
            return nil, nil
        } else {
            return nil, err
        }
    }
    return stat, nil
}

func FileExists(path string) (exists bool, err error) {
    stat, err := StatIfExists(path)
    if err != nil {
        return false, err
    }
    return stat != nil, nil
}

func IsDirectory(path string) (exists bool, err error) {
    stat, err := StatIfExists(path)
    if err != nil {
        return false, err
    }
    if stat == nil {
        return false, nil
    }
    return stat.IsDir(), nil
}

func ListDirectory(dirname string) (files []os.FileInfo, err error) {
    f, err := os.Open(dirname)
    if err != nil {
        log.Printf("Cannot open dir %s\n", dirname)
        return nil, err
    }
    defer f.Close()
    files, err = f.Readdir(-1)
    if err != nil {
        log.Printf("Cannot read dir %s\n", dirname)
        return nil, err
    }

    return files, nil
}

// Lists a directory; unlike ListDirectory this returns just the names, and sorts them
func ListDirectoryNames(dirname string) (files []string, err error) {
    f, err := os.Open(dirname)
    if err != nil {
        log.Printf("Cannot open dir %s\n", dirname)
        return nil, err
    }
    defer f.Close()
    files, err = f.Readdirnames(-1)
    if err != nil {
        log.Printf("Cannot read dir %s\n", dirname)
        return nil, err
    }

    sort.Strings(files)

    return files, nil
}

func TryReadTextFile(path string, defaultValue string) (string, error) {
    contents, err := ioutil.ReadFile(path)
    if err != nil {
        if patherr, ok := err.(*os.PathError); ok {
            if syserr, ok := patherr.Err.(syscall.Errno); ok {
                if syserr == 2 {
                    return defaultValue, nil
                }

                log.Printf("Error reading file.  code=%v\n", int(syserr))
            } else {
                log.Printf("Error reading file %T\n", patherr.Err)
            }
        } else {
            log.Printf("Error reading file %T\n", err)
        }
        return "", err
    }
    return string(contents), nil
}

func TryReadFile(path string) ([]byte, error) {
    contents, err := ioutil.ReadFile(path)
    if err != nil {
        if patherr, ok := err.(*os.PathError); ok {
            if syserr, ok := patherr.Err.(syscall.Errno); ok {
                if syserr == 2 {
                    return nil, nil
                }

                log.Printf("Error reading file.  code=%v\n", int(syserr))
            } else {
                log.Printf("Error reading file %T\n", patherr.Err)
            }
        } else {
            log.Printf("Error reading file %T\n", err)
        }
        return nil, err
    }
    return contents, nil
}
