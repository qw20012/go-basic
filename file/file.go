package file

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

func exists(name string, dir bool) bool {
	info, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}

	if dir {
		return info.IsDir()
	}

	return !info.IsDir()
}

// Identify whether given file exists
func IsFile(name string) bool {
	return exists(name, false)
}

// Identify whether given directory exists
func IsDir(dir string) bool {
	return exists(dir, true)
}

type FileFilters struct {
	OlderThan    time.Duration
	Prefix       string
	Suffix       string
	RegexPattern string
	CustomCheck  func(filename string) bool
	Callback     func(filename string) error
}

func DeleteFilesOlderThan(folder string, filter FileFilters) error {
	startScan := time.Now()
	return filepath.WalkDir(folder, func(osPathname string, de fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if osPathname == "" {
			return nil
		}
		if de.IsDir() {
			return nil
		}
		fileInfo, err := os.Stat(osPathname)
		if err != nil {
			return nil
		}
		fileName := fileInfo.Name()
		if filter.Prefix != "" && !strings.HasPrefix(fileName, filter.Prefix) {
			return nil
		}
		if filter.Suffix != "" && !strings.HasSuffix(fileName, filter.Suffix) {
			return nil
		}
		if filter.RegexPattern != "" {
			regex, err := regexp.Compile(filter.RegexPattern)
			if err != nil {
				return err
			}
			if !regex.MatchString(fileName) {
				return nil
			}
		}
		if filter.CustomCheck != nil && !filter.CustomCheck(osPathname) {
			return nil
		}
		if fileInfo.ModTime().Add(filter.OlderThan).Before(startScan) {
			if filter.Callback != nil {
				return filter.Callback(osPathname)
			} else {
				os.RemoveAll(osPathname)
			}
		}
		return nil
	},
	)
}

// DownloadFile to specified path
func DownloadFile(filepath string, url string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

// ReadFileWithReader and stream on a channel
func ReadFileWithReader(r io.Reader) (chan string, error) {
	out := make(chan string)
	go func() {
		defer close(out)
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			out <- scanner.Text()
		}
	}()

	return out, nil
}

// ReadFileWithReader with specific buffer size and stream on a channel
func ReadFileWithReaderAndBufferSize(r io.Reader, maxCapacity int) (chan string, error) {
	out := make(chan string)
	go func() {
		defer close(out)
		scanner := bufio.NewScanner(r)
		buf := make([]byte, maxCapacity)
		scanner.Buffer(buf, maxCapacity)
		for scanner.Scan() {
			out <- scanner.Text()
		}
	}()

	return out, nil
}

// ReadFile with filename
func ReadFile(filename string) (chan string, error) {
	if !IsFile(filename) {
		return nil, errors.New("file doesn't exist")
	}
	out := make(chan string)
	go func() {
		defer close(out)
		f, err := os.Open(filename)
		if err != nil {
			return
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			out <- scanner.Text()
		}
	}()

	return out, nil
}

// // ReadFile with filename and specific buffer size
func ReadFileWithBufferSize(filename string, maxCapacity int) (chan string, error) {
	if !IsFile(filename) {
		return nil, errors.New("file doesn't exist")
	}
	out := make(chan string)
	go func() {
		defer close(out)
		f, err := os.Open(filename)
		if err != nil {
			return
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		buf := make([]byte, maxCapacity)
		scanner.Buffer(buf, maxCapacity)
		for scanner.Scan() {
			out <- scanner.Text()
		}
	}()

	return out, nil
}

// CopyFile copies a file from source to dest. Any existing file will be overwritten
// and attributes will not be copied
func CopyFile(src string, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("can't stat %s: %w", src, err)
	}

	if !srcInfo.Mode().IsRegular() {
		return fmt.Errorf("can't copy non-regular source file %s (%s)", src, srcInfo.Mode().String())
	}

	srcFh, err := os.Open(src) //nolint
	if err != nil {
		return fmt.Errorf("can't open source file %s: %w", src, err)
	}
	defer srcFh.Close() //nolint

	err = os.MkdirAll(filepath.Dir(dst), 0750)
	if err != nil {
		return fmt.Errorf("can't make destination directory %s: %w", filepath.Dir(dst), err)
	}

	dstFh, err := os.Create(dst) //nolint
	if err != nil {
		return fmt.Errorf("can't create destination file %s: %w", dst, err)
	}
	defer dstFh.Close() //nolint

	size, err := io.Copy(dstFh, srcFh)
	if err != nil {
		return fmt.Errorf("can't copy data: %w", err)
	}
	if size != srcInfo.Size() {
		return fmt.Errorf("incomplete copy, %d of %d", size, srcInfo.Size())
	}
	return dstFh.Sync()
}

// CopyDir copies all files from src to dst, recursively
func CopyDir(src string, dst string) error {
	list, err := ListFiles(src)
	if err != nil {
		return fmt.Errorf("can't list source files in %s: %w", src, err)
	}
	for _, srcFile := range list {
		stripSrcDir := strings.TrimPrefix(srcFile, src)
		dstFile := filepath.Join(dst, stripSrcDir)
		if err = CopyFile(srcFile, dstFile); err != nil {
			return fmt.Errorf("can't copy %s to %s: %w", srcFile, dstFile, err)
		}
	}
	return nil
}

// ListFiles gets recursive list of all files in a directory
func ListFiles(directory string) (list []string, err error) {
	err = filepath.Walk(directory, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}
		if info.IsDir() {
			return nil
		}
		list = append(list, path)
		return nil
	})
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	return list, err
}

// TempFileName returns a new temporary file name in the directory dir.
// The filename is generated by taking pattern and adding a random
// string to the end. If pattern includes a "*", the random string
// replaces the last "*".
// If dir is the empty string, TempFileName uses the default directory
// for temporary files (see os.TempDir).
// Multiple programs calling TempFileName simultaneously
// will not choose the same file name.
// some code borrowed from stdlib https://golang.org/src/io/ioutil/tempfile.go
func TempFileName(dir, pattern string) (string, error) {
	var once sync.Once
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})
	// prefixAndSuffix splits pattern by the last wildcard "*", if applicable,
	// returning prefix as the part before "*" and suffix as the part after "*".
	prefixAndSuffix := func(pattern string) (prefix, suffix string) {
		if pos := strings.LastIndex(pattern, "*"); pos != -1 {
			prefix, suffix = pattern[:pos], pattern[pos+1:]
		} else {
			prefix = pattern
		}
		return
	}

	if dir == "" {
		dir = os.TempDir()
	}

	prefix, suffix := prefixAndSuffix(pattern)

	for i := 0; i < 10000; i++ {
		name := filepath.Join(dir, prefix+fmt.Sprintf("%x", rand.Int())+suffix) //nolint
		_, err := os.Stat(name)
		if os.IsNotExist(err) {
			return name, nil
		}
	}
	return "", errors.New("can't generate temp file name")
}

// SanitizePath removes invalid characters from path
func SanitizePath(s string) string {
	return regexp.MustCompile(`[<>:"|?*]+`).ReplaceAllString(filepath.Clean(s), "_")
}

func resolve(name string) string {
	// This implementation is based on Dir.Open's code in the standard net/http package.
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
		strings.Contains(name, "\x00") {
		return ""
	}

	dir := name
	if dir == "" {
		dir = "."
	}

	return filepath.Join(dir, filepath.FromSlash(filepath.Clean(name)))
}

// Copy copies a file or directory from src to dst. If it is
// a directory, all of the files and sub-directories will be copied.
func Copy(src, dst string) error {
	if src = resolve(src); src == "" {
		return os.ErrNotExist
	}

	if dst = resolve(dst); dst == "" {
		return os.ErrNotExist
	}

	if dst == src {
		return os.ErrInvalid
	}

	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return CopyDir(src, dst)
	}

	return CopyFile(src, dst)
}
