package loggist

// log level
// the lefter tag has the higher level
// to classify your log level
// FALAL > ERROR > WARN > DEBUG > INFO
const (
	DEBUG_LOG_KEY = "DEBUG"
	ERROR_LOG_KEY = "ERROR"
	FATAL_LOG_KEY = "FATAL"
	INFO_LOG_KEY  = "INFO"
	WARN_LOG_KEY  = "WARN"
)

type LogLevel struct {
	level  string
	prefix string
}

// get log level name
func (self *LogLevel) getLevelName() string {
	return self.level
}

// get log level prefix
func (self *LogLevel) getLevelPrefix() string {
	return self.prefix
}

var logLevelMap map[string]LogLevel = map[string]LogLevel{
	INFO_LOG_KEY:  LogLevel{level: INFO_LOG_KEY, prefix: "[" + INFO_LOG_KEY + "]"},
	DEBUG_LOG_KEY: LogLevel{level: DEBUG_LOG_KEY, prefix: "[" + DEBUG_LOG_KEY + "]"},
	WARN_LOG_KEY:  LogLevel{level: WARN_LOG_KEY, prefix: "[" + WARN_LOG_KEY + "]"},
	ERROR_LOG_KEY: LogLevel{level: ERROR_LOG_KEY, prefix: "[" + ERROR_LOG_KEY + "]"},
	FATAL_LOG_KEY: LogLevel{level: FATAL_LOG_KEY, prefix: "[" + FATAL_LOG_KEY + "]"},
}

// getInfoLogTag returns the info level map structure
func getInfoLogTag() LogLevel {
	return logLevelMap[INFO_LOG_KEY]
}

// getDebugLogTag returns the debug level map structure
func getDebugLogTag() LogLevel {
	return logLevelMap[DEBUG_LOG_KEY]
}

// getWarnLogTag returns the warn level map structure
func getWarnLogTag() LogLevel {
	return logLevelMap[WARN_LOG_KEY]
}

// getErrorLogTag returns the error level map structure
func getErrorLogTag() LogLevel {
	return logLevelMap[ERROR_LOG_KEY]
}

// getFatalLogTag returns the fatal level map structure
func getFatalLogTag() LogLevel {
	return logLevelMap[FATAL_LOG_KEY]
}
