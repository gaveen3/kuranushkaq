package website

import (
	"os"
	"path/filepath"
	"runtime"

	log "github.com/Sirupsen/logrus"
)

func init() {

	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)

	if pc, file, line, ok := runtime.Caller(2); ok {
		fName := runtime.FuncForPC(pc).Name()
		log.WithFields(log.Fields{
			"file": filepath.Base(file),
			"line": line,
			"func": fName,
		})
	}

	switch getENV("LOG_LEVEL") {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
}

//LogInfo *
func LogInfo(v ...interface{}) {
	log.Info(v...)
}

//LogInfoln *
func LogInfoln(v ...interface{}) {
	log.Infoln(v...)
}

//LogInfof *
func LogInfof(format string, v ...interface{}) {
	log.Infof(format, v...)
}

//LogDebug *
func LogDebug(v ...interface{}) {
	log.Debug(v...)
}

//LogDebugln *
func LogDebugln(v ...interface{}) {
	log.Debugln(v...)
}

//LogDebugf *
func LogDebugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

//LogWarn *
func LogWarn(v ...interface{}) {
	log.Warn(v...)
}

//LogWarnln *
func LogWarnln(v ...interface{}) {
	log.Warnln(v...)
}

//LogWarnf *
func LogWarnf(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

//LogError *
func LogError(v ...interface{}) {
	log.Error(v...)
}

//LogErrorln *
func LogErrorln(v ...interface{}) {
	log.Errorln(v...)
}

//LogErrorf *
func LogErrorf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

//LogFatal *
func LogFatal(v ...interface{}) {
	log.Fatal(v...)
}

//LogFatalln *
func LogFatalln(v ...interface{}) {
	log.Fatalln(v...)
}

//LogFatalf *
func LogFatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

//LogPanic *
func LogPanic(v ...interface{}) {
	log.Panic(v...)
}

//LogPanicln *
func LogPanicln(v ...interface{}) {
	log.Panicln(v...)
}

//LogPanicf *
func LogPanicf(format string, v ...interface{}) {
	log.Panicf(format, v...)
}
