package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry有一个用于存放目录绝对路径的字段
type DirEntry struct {
	absDir string
}

// newDirEntry（）先把参数转换成绝对路径，若转换出错，则panic终止程序，否则创建DirEntry实例并返回。
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

// readClass（）先把目录和class文件拼成一条完整路径，调用ioutil包的ReadFile（）函数读取class文件内容，最后返回。String()函数直接返回目录。
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
