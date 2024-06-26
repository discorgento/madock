package paths

import (
	"github.com/faradey/madock/src/helper/hash"
	"github.com/faradey/madock/src/helper/logger"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func GetExecDirPath() string {
	var dirAbsPath string

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exReal, err := filepath.EvalSymlinks(ex)
	if err != nil {
		dirAbsPath = filepath.Dir(ex)
		return dirAbsPath
	} else {
		dirAbsPath = filepath.Dir(exReal)
		return dirAbsPath
	}

	panic("Unknown error")
}

func GetExecDirName() string {
	return filepath.Base(GetExecDirPath())
}

func GetExecDirNameByPath(path string) string {
	return filepath.Base(path)
}

func GetRunDirPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}

	return dir
}

func GetRunDirName() string {
	return filepath.Base(GetRunDirPath())
}

func GetRunDirNameWithHash() string {
	return filepath.Base(GetRunDirPath()) + "__" + strconv.Itoa(int(hash.Hash(GetRunDirPath())))
}

func GetDirs(path string) (dirs []string) {
	items, err := os.ReadDir(path)
	if err != nil {
		logger.Fatal(err)
	}

	for _, file := range items {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}

	return dirs
}

func GetFiles(path string) (dirs []string) {
	items, err := os.ReadDir(path)
	if err != nil {
		logger.Fatal(err)
	}

	for _, file := range items {
		if !file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}

	return dirs
}

func GetFilesRecursively(path string) (dirs []string) {
	items, err := os.ReadDir(path)
	if err == nil {
		for _, file := range items {
			if !file.IsDir() {
				dirs = append(dirs, path+"/"+file.Name())
			} else {
				dirs = append(dirs, GetFilesRecursively(path+"/"+file.Name())...)
			}
		}
	}

	return dirs
}

func GetDBFiles(path string) (dirs []string) {
	items, err := os.ReadDir(path)
	if err != nil {
		logger.Fatal(err)
	}

	for _, file := range items {
		if !file.IsDir() {
			if file.Name()[0:1] != "." &&
				strings.Contains(strings.ToLower(file.Name()), ".sql") &&
				!strings.Contains(strings.ToLower(path), "/dev/tests/acceptance") &&
				!strings.Contains(strings.ToLower(path), strings.ToLower(strings.Trim(GetRunDirPath(), "/"))+"/vendor/") {
				dirs = append(dirs, path+"/"+file.Name())
			}
		} else {
			dirs = append(dirs, GetDBFiles(path+"/"+file.Name())...)
		}
	}

	return dirs
}

func MakeDirsByPath(val string) string {
	trimVal := strings.Trim(val, "/")
	if trimVal != "" {
		dirs := strings.Split(trimVal, "/")
		var err error
		for i := 0; i < len(dirs); i++ {
			if !IsFileExist("/" + strings.Join(dirs[:i+1], "/")) {
				err = os.Mkdir("/"+strings.Join(dirs[:i+1], "/"), 0755)
				if err != nil {
					logger.Fatal(err)
				}
			}
		}
	}

	return val
}

func GetActiveProjects() []string {
	var activeProjects []string
	cmd := exec.Command("docker", "ps", "--format", "json")
	result, err := cmd.CombinedOutput()
	if err != nil {
		logger.Println(err, string(result))
	} else {
		resultString := string(result)
		projects := GetDirs(MakeDirsByPath(GetExecDirPath() + "/aruntime/projects"))
		for _, projectName := range projects {
			if strings.Contains(resultString, strings.ToLower(projectName)+"-") {
				activeProjects = append(activeProjects, projectName)
			}
		}
	}

	return activeProjects
}

func IsFileExist(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}

	return false
}

func Copy(srcpath, dstpath string) (err error) {
	r, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer r.Close() // ignore error: file was opened read-only.

	w, err := os.Create(dstpath)
	if err != nil {
		return err
	}

	defer func() {
		// Report the error, if any, from Close, but do so
		// only if there isn't already an outgoing error.
		if c := w.Close(); err == nil {
			err = c
		}
	}()

	_, err = io.Copy(w, r)
	return err
}
