package loggist

import (
	"fmt"
	"log"
	"os"
	"time"
)

// log mode
type LogMode int

// different log mode
const (
	MODE_CONSOLE LogMode = iota // write log to console for debug
	MODE_FILE                   // write log to file
)

// log record mode
// only used under the file log mode
type LogRecordMode int

// different log record mode
// the order of the constant below is corresponding to the conf.xml file
const (
	RECORD_MODE_YM    LogRecordMode = iota // Logs/2016/201612
	RECORD_MODE_YMD                        // Logs/201612/20161211
	RECORD_MODE_YMDH                       // Logs/20161228/20161228-16
	RECORD_MODE_YMDHM                      // Logs/20161228-16/20161228-16-40
)

// instantiate logger
func NewLoggist(mode LogMode, recordMode LogRecordMode) *Loggist {
	logger := &Loggist{}
	logger.setDefaultPosition()
	logger.setLogMode(mode)
	logger.setLogRecordMode(recordMode)
	logger.logger = &log.Logger{}
	logger.conf = &LogConf{}
	logger.conf.initLogConf()
	return logger
}

// Loggist defination
type Loggist struct {
	posision   int
	logger     *log.Logger
	mode       LogMode
	recordMode LogRecordMode
	conf       *LogConf
}

// return the current invoke position
func (self *Loggist) getPosition() int {
	return self.posision
}

// set the invoke position
func (self *Loggist) setPosition(pos int) {
	self.posision = pos
}

// set the default invoke position
func (self *Loggist) setDefaultPosition() {
	self.posision = 1
}

// set log mode
func (self *Loggist) setLogMode(mode LogMode) {
	if mode != MODE_FILE && mode != MODE_CONSOLE {
		panic("Undefine Log mode")
	}
	self.mode = mode
}

// return the currunt log mode
func (self *Loggist) getLogMode() LogMode {
	return self.mode
}

// set log record mode
func (self *Loggist) setLogRecordMode(mode LogRecordMode) {
	if mode != RECORD_MODE_YM &&
		mode != RECORD_MODE_YMD &&
		mode != RECORD_MODE_YMDH &&
		mode != RECORD_MODE_YMDHM {
		panic("Undefined Log Record Mode")
	}
	self.recordMode = mode
}

// return the current log record mode
func (self *Loggist) getLogRecordMode() LogRecordMode {
	if self.recordMode == 0 {
		self.recordMode = RECORD_MODE_YM
	}
	return self.recordMode
}

// append log content at the end of the file
// if the file doesn't exist, then create it
func (self *Loggist) write(content string) {
	self.logger.SetFlags(log.LstdFlags)
	// get log mode
	mode := self.getLogMode()
	// different mode for different dealway
	switch mode {
	case MODE_FILE:
		self.dealFileMode(content)
	case MODE_CONSOLE:
		self.dealConsoleMode(content)
	default:
		panic("Undefine Log Mode")
	}
}

// log console mode
func (self *Loggist) dealConsoleMode(content string) {
	self.logger.SetOutput(os.Stdout)
	self.logger.Print(content)
}

// log file mode
func (self *Loggist) dealFileMode(content string) {
	fileName := self.genFullFileName()
	fd, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0644)
	if err != nil {
		if os.IsPermission(err) {
			panic(err)
		}
		if os.IsNotExist(err) {
			mkDir(fileName)
		}
		fd, err = os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_SYNC, 0664)
		if err != nil {
			panic(err)
		}
	}
	self.logger.SetOutput(fd)
	self.logger.Print(content)
	defer fd.Close()
}

// generate log file name and log store path as full file Name
func (self *Loggist) genFullFileName() string {
	var fileName string
	mode := self.getLogRecordMode()
	conf := self.conf
	rootDir := conf.LogRoot
	t := time.Now()
	year, month, day := t.Date()
	hour := t.Hour()
	min := t.Minute()
	subDirFmt := conf.LogFormat.Mode[mode].SubLogDir
	fileNameFmt := conf.LogFormat.Mode[mode].FileFmt
	pathSep := string(os.PathSeparator)
	switch mode {
	case RECORD_MODE_YM:
		fileName = rootDir + pathSep +
			fmt.Sprintf(subDirFmt, year) + pathSep +
			fmt.Sprintf(fileNameFmt, year, month)
	case RECORD_MODE_YMD:
		fileName = rootDir + pathSep +
			fmt.Sprintf(subDirFmt, year, month) + pathSep +
			fmt.Sprintf(fileNameFmt, year, month, day)
	case RECORD_MODE_YMDH:
		fileName = rootDir + pathSep +
			fmt.Sprintf(subDirFmt, year, month, day) + pathSep +
			fmt.Sprintf(fileNameFmt, year, month, day, hour)
	case RECORD_MODE_YMDHM:
		fileName = rootDir + pathSep +
			fmt.Sprintf(subDirFmt, year, month, day, hour) + pathSep +
			fmt.Sprintf(fileNameFmt, year, month, day, hour, min)
	default:
		panic("Undefined Log Record Mode")

	}
	return fileName
}

// logger.Info(d1, d2, d3)
func (self *Loggist) Info(v ...interface{}) {
	content := generateLogContent(getInfoLogTag(), self.getPosition(), "", v...)
	self.write(content)
}

// logger.Infof(format, d1, d2, d3)
func (self *Loggist) Infof(format string, v ...interface{}) {
	content := generateLogContent(getInfoLogTag(), self.getPosition(), format, v...)
	self.write(content)
}

// logger.Debug(d1, d2, d3)
func (self *Loggist) Debug(v ...interface{}) {
	content := generateLogContent(getDebugLogTag(), self.getPosition(), "", v...)
	self.write(content)
}

// logger.Debugf(format, d1, d2, d3)
func (self *Loggist) Debugf(format string, v ...interface{}) {
	content := generateLogContent(getDebugLogTag(), self.getPosition(), format, v...)
	self.write(content)
}

// logger.Warn(d1, d2, d3)
func (self *Loggist) Warn(v ...interface{}) {
	content := generateLogContent(getWarnLogTag(), self.getPosition(), "", v...)
	self.write(content)
}

// logger.Warnf(format, d1, d2, d3)
func (self *Loggist) Warnf(format string, v ...interface{}) {
	content := generateLogContent(getWarnLogTag(), self.getPosition(), format, v...)
	self.write(content)
}

// logger.Error(d1, d2, d3)
func (self *Loggist) Error(v ...interface{}) {
	content := generateLogContent(getErrorLogTag(), self.getPosition(), "", v...)
	self.write(content)
}

// logger.Errorf(format, d1, d2, d3)
func (self *Loggist) Errorf(format string, v ...interface{}) {
	content := generateLogContent(getErrorLogTag(), self.getPosition(), format, v...)
	self.write(content)
}

// logger.Fatal(d1, d2, d3)
func (self *Loggist) Fatal(v ...interface{}) {
	content := generateLogContent(getFatalLogTag(), self.getPosition(), "", v...)
	self.write(content)
}

// logger.Fatalf(format, d1, d2, d3)
func (self *Loggist) Fatalf(format string, v ...interface{}) {
	content := generateLogContent(getFatalLogTag(), self.getPosition(), format, v...)
	self.write(content)
}
