package loggist

import (
	"testing"
	"time"
)

// define test level for testing
var levelKeys []string = []string{INFO_LOG_KEY, DEBUG_LOG_KEY, WARN_LOG_KEY, ERROR_LOG_KEY, FATAL_LOG_KEY}

// log level test
func TestLevel(t *testing.T) {
	var tmpTag LogLevel
	for _, v := range levelKeys {
		switch v {
		case INFO_LOG_KEY:
			tmpTag = getInfoLogTag()
			outputTagInfo(tmpTag, t)
		case DEBUG_LOG_KEY:
			tmpTag = getDebugLogTag()
			outputTagInfo(tmpTag, t)
		case WARN_LOG_KEY:
			tmpTag = getWarnLogTag()
			outputTagInfo(tmpTag, t)
		case ERROR_LOG_KEY:
			tmpTag = getErrorLogTag()
			outputTagInfo(tmpTag, t)
		case FATAL_LOG_KEY:
			tmpTag = getFatalLogTag()
			outputTagInfo(tmpTag, t)
		default:
			t.Fatalf("Unknow log level: lelel = %s, prefix = %s", tmpTag.getLevelName(), tmpTag.getLevelPrefix())
		}
	}
}

func outputTagInfo(tmp LogLevel, t *testing.T) {
	t.Logf("level = %s prefix = %s", tmp.getLevelName(), tmp.getLevelPrefix())
}

// test holding the same file handler
func TestFileHandler(t *testing.T) {
	var format string
	var contentString string
	var contentNum int
	fl := NewLoggist(MODE_FILE, RECORD_MODE_YMDHM)

	var ticker *time.Ticker = time.NewTicker(1 * time.Second)
	go func() {
		for t := range ticker.C {
			format = "string=%s,num=%d,time=%v"
			contentString = "Infof"
			contentNum = 12345
			fl.Infof(format, contentString, contentNum, t)
		}
	}()
	time.Sleep(time.Minute * 5) //阻塞，则执行次数为sleep的休眠时间/ticker的时间
	ticker.Stop()
}

// different log level test
func TestLevelWrite(t *testing.T) {
	var format string
	var contentString string
	var contentNum int
	fl := NewLoggist(MODE_FILE, RECORD_MODE_YMDH)
	// Info
	format = ""
	contentString = "Info"
	contentNum = 12345
	fl.Info(contentString, contentNum)
	// Infof
	format = "string=%s,num=%d"
	contentString = "Infof"
	contentNum = 12345
	fl.Infof(format, contentString, contentNum)

	// Debug
	format = ""
	contentString = "Debug"
	contentNum = 12345
	fl.Debug(contentString, contentNum)
	// Debugf
	format = "string=%s,num=%d"
	contentString = "Debugf"
	contentNum = 12345
	fl.Debugf(format, contentString, contentNum)

	// Warn
	format = ""
	contentString = "Warn"
	contentNum = 12345
	fl.Warn(contentString, contentNum)
	// Errorf
	format = "string=%s,num=%d"
	contentString = "Warnf"
	contentNum = 12345
	fl.Warnf(format, contentString, contentNum)

	// Error
	format = ""
	contentString = "Error"
	contentNum = 12345
	fl.Error(contentString, contentNum)
	// Errorf
	format = "string=%s,num=%d"
	contentString = "Errorf"
	contentNum = 12345
	fl.Errorf(format, contentString, contentNum)

	// Fatal
	format = ""
	contentString = "Fatal"
	contentNum = 12345
	fl.Fatal(contentString, contentNum)
	// Fatalf
	format = "string=%s,num=%d"
	contentString = "Fatalf"
	contentNum = 12345
	fl.Fatalf(format, contentString, contentNum)
}

// test record mode
func TestRecordMode(t *testing.T) {
	var format string
	var contentString string
	var contentNum int
	var fl *Loggist
	var rmode LogRecordMode
	fl = NewLoggist(MODE_FILE, RECORD_MODE_YM)
	// YM
	rmode = fl.getLogRecordMode()
	t.Log(rmode)
	contentString = "Info"
	contentNum = 12345
	fl.Info(contentString, contentNum)
	format = "string=%s,num=%d"
	contentString = "Infof"
	contentNum = 12345
	fl.Infof(format, contentString, contentNum)

	// YMD
	fl.setLogRecordMode(RECORD_MODE_YMD)
	rmode = fl.getLogRecordMode()
	t.Log(rmode)
	contentString = "Info"
	contentNum = 12345
	fl.Info(contentString, contentNum)
	format = "string=%s,num=%d"
	contentString = "Infof"
	contentNum = 12345
	fl.Infof(format, contentString, contentNum)

	// YMDH
	fl.setLogRecordMode(RECORD_MODE_YMDH)
	rmode = fl.getLogRecordMode()
	t.Log(rmode)
	contentString = "Info"
	contentNum = 12345
	fl.Info(contentString, contentNum)
	format = "string=%s,num=%d"
	contentString = "Infof"
	contentNum = 12345
	fl.Infof(format, contentString, contentNum)

	// YMDHM
	fl.setLogRecordMode(RECORD_MODE_YMDHM)
	rmode = fl.getLogRecordMode()
	t.Log(rmode)
	contentString = "Info"
	contentNum = 12345
	fl.Info(contentString, contentNum)
	format = "string=%s,num=%d"
	contentString = "Infof"
	contentNum = 12345
	fl.Infof(format, contentString, contentNum)
}

// output mode test
func TestOutputMode(t *testing.T) {
	var format string
	var contentString string
	var contentNum int
	var fl *Loggist
	// File Mode Test
	fl = NewLoggist(MODE_FILE, RECORD_MODE_YMD)
	format = ""
	contentString = "Info"
	contentNum = 12345
	fl.Info(contentString, contentNum)

	format = "string=%s,num=%d"
	contentString = "Infof"
	contentNum = 12345
	fl.Infof(format, contentString, contentNum)

	// console Mode Test
	fl = NewLoggist(MODE_CONSOLE, RECORD_MODE_YMD)
	format = ""
	contentString = "Debug"
	contentNum = 12345
	fl.Debug(contentString, contentNum)

	format = "string=%s,num=%d"
	contentString = "Debugf"
	contentNum = 12345
	fl.Debugf(format, contentString, contentNum)
}
