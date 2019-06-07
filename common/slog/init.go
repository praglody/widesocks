package slog

import "log"

const (
	LOG_DEBUG uint8 = 1 << iota
	LOG_VERBOSE
	LOG_INFO
	LOG_NOTICE
	LOG_WARNING
	LOG_ERROR
	LOG_CRITICAL
	LOG_ALERT
	// LOG_EMERGENCY
)

var logLevel = LOG_INFO

var logLevelDesc = map[uint8]string{
	LOG_DEBUG:    "DEBUG",
	LOG_VERBOSE:  "VERBOSE",
	LOG_INFO:     "INFO",
	LOG_NOTICE:   "NOTICE",
	LOG_WARNING:  "WARNING",
	LOG_ERROR:    "ERROR",
	LOG_CRITICAL: "CRITICAL",
	LOG_ALERT:    "ALERT",
}

func SetLogLevel(level uint8) {
	if logLevelDesc[level] != "" {
		logLevel = level
		log.Printf("Set the LOG LEVEL to %s", logLevelDesc[level])
	} else {
		Warning("Set Log Level Error!")
	}
}

func Debug(v ...interface{}) {
	if logLevel <= LOG_DEBUG {
		log.Println(v)
	}
}

func Debugf(format string, v ...interface{}) {
	if logLevel <= LOG_DEBUG {
		log.Printf(format, v)
	}
}

func Verbose(v ...interface{}) {
	if logLevel <= LOG_VERBOSE {
		log.Println(v)
	}
}

func Verbosef(format string, v ...interface{}) {
	if logLevel <= LOG_VERBOSE {
		log.Printf(format, v)
	}
}

func Info(v ...interface{}) {
	if logLevel <= LOG_INFO {
		log.Println(v)
	}
}

func Infof(format string, v ...interface{}) {
	if logLevel <= LOG_INFO {
		log.Printf(format, v)
	}
}

func Notice(v ...interface{}) {
	if logLevel <= LOG_NOTICE {
		log.Println(v)
	}
}

func Noticef(format string, v ...interface{}) {
	if logLevel <= LOG_NOTICE {
		log.Printf(format, v)
	}
}

func Warning(v ...interface{}) {
	if logLevel <= LOG_INFO {
		log.Println(v)
	}
}

func Warningf(format string, v ...interface{}) {
	if logLevel <= LOG_INFO {
		log.Printf(format, v)
	}
}

func Error(v ...interface{}) {
	if logLevel <= LOG_ERROR {
		log.Println(v)
	}
}

func Errorf(format string, v ...interface{}) {
	if logLevel <= LOG_ERROR {
		log.Println(v)
	}
}

func Emergency(v ...interface{}) {
	log.Panicln(v)
}

func Emergencyf(format string, v ...interface{}) {
	log.Panicf(format, v)
}
