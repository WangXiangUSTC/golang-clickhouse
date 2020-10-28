package clickhouse

type logger struct {
	debug func(message string)
	info  func(message string)
	warn  func(message string)
	error func(message string)
	fatal func(message string)
}

// Debug sets logger for debug
func Debug(callback func(message string)) {
	cfg.logger.debug = callback
	cfg.logger.debug("Set custom debug logger")
}

// Info sets logger for into
func Info(callback func(message string)) {
	cfg.logger.info = callback
	cfg.logger.debug("Set custom info logger")
}

// Warn sets logger for warning
func Warn(callback func(message string)) {
	cfg.logger.warn = callback
	cfg.logger.debug("Set custom warning logger")
}

// Error sets logger for error
func Error(callback func(message string)) {
	cfg.logger.error = callback
	cfg.logger.debug("Set custom error logger")
}

// Fatal sets logger for fatal
func Fatal(callback func(message string)) {
	cfg.logger.fatal = callback
	cfg.logger.debug("Set custom fatal logger")
}