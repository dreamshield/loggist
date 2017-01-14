package loggist

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Logger interface {
	Info(v ...interface{}) string
	Infof(format string, v ...interface{}) string
	Debug(v ...interface{}) string
	Debugf(format string, v ...interface{}) string
	Warn(v ...interface{}) string
	Warnf(format string, v ...interface{}) string
	Error(v ...interface{}) string
	Errorf(format string, v ...interface{}) string
	Fatal(v ...interface{}) string
	Fatalf(format string, v ...interface{}) string
}

// make directory
func mkDir(fileName string) {
	paths := strings.Split(fileName, string(os.PathSeparator))
	paths = paths[0 : len(paths)-1]
	if err := os.MkdirAll(strings.Join(paths, string(os.PathSeparator)), 0777); err != nil {
		panic(err)
	}
}

// create file
func createFile(fileName string) (fd *os.File, err error) {
	fd, err = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0644)
	if err != nil {
		if os.IsPermission(err) {
			return nil, err
		}
		if os.IsNotExist(err) {
			mkDir(fileName)
		}
		fd, err = os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_SYNC, 0664)
		if err != nil {
			return nil, err
		}
	}
	return fd, err
}

// generate log content
func generateLogContent(level LogLevel, pos int, format string, v ...interface{}) string {
	deepNum := int(pos) + 2
	baseInfo := fmt.Sprintf("%s %s - ", level.getLevelPrefix(), getInvokePosition(deepNum))
	var result string
	if len(format) > 0 {
		result = fmt.Sprintf((baseInfo + format), v...)
	} else {
		vLen := len(v)
		params := make([]interface{}, (vLen + 1))
		params[0] = baseInfo
		for i := 1; i <= vLen; i++ {
			params[i] = v[i-1]
		}
		result = fmt.Sprint(params...)
	}
	return result
}

// get the log trace level
func getInvokePosition(deepNum int) string {
	pc, file, line, ok := runtime.Caller(deepNum)
	if !ok {
		return ""
	}
	fileName := ""
	if index := strings.LastIndex(file, "/"); index > 0 {
		fileName = file[index+1 : len(file)]
	}
	funcPath := ""
	funcPtr := runtime.FuncForPC(pc)
	if funcPtr != nil {
		funcPath = funcPtr.Name()
	}
	return fmt.Sprintf("%s - %s:%d", funcPath, fileName, line)
}
