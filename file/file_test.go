package file

import (
	"testing"
)

func TestIsFile(t *testing.T) {
	tests := map[string]bool{
		"file.go": true,
		"aaa.bbb": false,
		"/":       false,
	}
	for fpath, mustExist := range tests {
		if IsFile(fpath) != mustExist {
			t.Fatalf("TestIsFile expected Type=%v, Got=%v", mustExist, IsFile(fpath))
		}
	}
}

func TestFolderExists(t *testing.T) {
	tests := map[string]bool{
		".":       true,
		"../file": true,
		"aabb":    false,
	}
	for fpath, mustExist := range tests {
		if IsDir(fpath) != mustExist {
			t.Fatalf("TestIsFile expected Type=%v, Got=%v", mustExist, IsDir(fpath))
		}
	}
}

/*
func TestDeleteFilesOlderThan(t *testing.T) {
	// create a temporary folder with a couple of files
	fo, err := os.MkdirTemp("", "")
	//require.Nil(t, err, "couldn't create folder: %s", err)
	ttl := time.Duration(5 * time.Second)

	// defer temporary folder removal
	defer os.RemoveAll(fo)
	checkFolderErr := func(err error) {
		//require.Nil(t, err, "couldn't create folder: %s", err)
	}
	checkFiles := func(fileName string) {
		//require.False(t, FileExists(fileName), "file \"%s\" still exists", fileName)
	}
	createFile := func() string {
		fi, _ := os.CreateTemp(fo, "")
		//require.Nil(t, err, "couldn't create f: %s", err)
		fName := fi.Name()
		fi.Close()
		return fName
	}
	t.Run("prefix props test", func(t *testing.T) {
		fName := createFile()
		fileInfo, _ := os.Stat(fName)
		// sleep for 5 seconds
		time.Sleep(5 * time.Second)
		// delete files older than 5 seconds
		filter := FileFilters{
			OlderThan: ttl,
			Prefix:    fileInfo.Name(),
		}
		err = DeleteFilesOlderThan(fo, filter)
		checkFolderErr(err)
		checkFiles(fName)
	})
	t.Run("suffix props test", func(t *testing.T) {
		fName := createFile()
		fileInfo, _ := os.Stat(fName)
		// sleep for 5 seconds
		time.Sleep(5 * time.Second)
		// delete files older than 5 seconds
		filter := FileFilters{
			OlderThan: ttl,
			Suffix:    string(fileInfo.Name()[len(fileInfo.Name())-1]),
		}
		err = DeleteFilesOlderThan(fo, filter)
		checkFolderErr(err)
		checkFiles(fName)
	})
	t.Run("regex pattern props test", func(t *testing.T) {
		fName := createFile()
		fName1 := createFile()

		// sleep for 5 seconds
		time.Sleep(5 * time.Second)
		// delete files older than 5 seconds
		filter := FileFilters{
			OlderThan:    ttl,
			RegexPattern: "[0-9]",
		}
		err = DeleteFilesOlderThan(fo, filter)
		checkFolderErr(err)
		checkFiles(fName)
		checkFiles(fName1)
	})
	t.Run("custom check props test", func(t *testing.T) {
		fName := createFile()
		fName1 := createFile()

		// sleep for 5 seconds
		time.Sleep(5 * time.Second)
		// delete files older than 5 seconds
		filter := FileFilters{
			OlderThan: ttl,
			CustomCheck: func(filename string) bool {
				return true
			},
		}
		err = DeleteFilesOlderThan(fo, filter)
		checkFolderErr(err)
		checkFiles(fName)
		checkFiles(fName1)
	})
	t.Run("custom check props negative test", func(t *testing.T) {
		//fName := createFile()
		// sleep for 5 seconds
		time.Sleep(5 * time.Second)
		// delete files older than 5 seconds
		filter := FileFilters{
			OlderThan: ttl,
			CustomCheck: func(filename string) bool {
				return false // should not delete the file
			},
		}
		err = DeleteFilesOlderThan(fo, filter)
		checkFolderErr(err)
		//require.True(t, FileExists(fName), "file \"%s\" should exists", fName)
	})
	t.Run("callback props test", func(t *testing.T) {
		fName := createFile()
		fName1 := createFile()

		// sleep for 5 seconds
		time.Sleep(5 * time.Second)
		// delete files older than 5 seconds
		filter := FileFilters{
			OlderThan: ttl,
			CustomCheck: func(filename string) bool {
				return true
			},
			Callback: func(filename string) error {
				t.Log("deleting file manually")
				return os.Remove(filename)
			},
		}
		err = DeleteFilesOlderThan(fo, filter)
		checkFolderErr(err)
		checkFiles(fName)
		checkFiles(fName1)
	})
}

func TestDownloadFile(t *testing.T) {
	// attempt to download http://ipv4.download.thinkbroadband.com/5MB.zip to temp folder
	tmpfile, err := os.CreateTemp("", "")

	//require.Nil(t, err, "couldn't create folder: %s", err)
	fname := tmpfile.Name()

	os.Remove(fname)

	err = DownloadFile(fname, "http://ipv4.download.thinkbroadband.com/5MB.zip")
	//require.Nil(t, err, "couldn't download file: %s", err)

	//require.True(t, FileExists(fname), "file \"%s\" doesn't exists", fname)

	// remove the downloaded file
	os.Remove(fname)
}

func TestReadFile(t *testing.T) {
	fileContent := `test
	test1
	test2`
	f, err := os.CreateTemp("", "")
	require.Nil(t, err, "couldn't create file: %s", err)
	fname := f.Name()
	_, _ = f.Write([]byte(fileContent))
	f.Close()
	defer os.Remove(fname)

	fileContentLines := strings.Split(fileContent, "\n")
	// compare file lines
	c, err := ReadFile(fname)
	require.Nil(t, err, "couldn't open file: %s", err)
	i := 0
	for line := range c {
		require.Equal(t, fileContentLines[i], line, "lines don't match")
		i++
	}
}

func TestReadFileWithBufferSize(t *testing.T) {
	fileContent := `test
	test1
	test2`
	f, err := os.CreateTemp("", "")
	require.Nil(t, err, "couldn't create file: %s", err)
	fname := f.Name()
	_, _ = f.Write([]byte(fileContent))
	f.Close()
	defer os.Remove(fname)

	fileContentLines := strings.Split(fileContent, "\n")
	// compare file lines
	c, err := ReadFileWithBufferSize(fname, 64*1024)
	require.Nil(t, err, "couldn't open file: %s", err)
	i := 0
	for line := range c {
		require.Equal(t, fileContentLines[i], line, "lines don't match")
		i++
	}
}
*/
