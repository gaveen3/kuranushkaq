package rclog

import (
	"os"
	"path/filepath"
	"runtime"

	"utils"

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

	switch utils.GetENV("LOG_LEVEL") {
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

//Info *
func Info(v ...interface{}) {
	log.Info(v...)
}

//Infoln *
func Infoln(v ...interface{}) {
	log.Infoln(v...)
}

//Infof *
func Infof(format string, v ...interface{}) {
	log.Infof(format, v...)
}

//Debug *
func Debug(v ...interface{}) {
	log.Debug(v...)
}

//Debugln *
func Debugln(v ...interface{}) {
	log.Debugln(v...)
}

//Debugf *
func Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

//Warn *
func Warn(v ...interface{}) {
	log.Warn(v...)
}

//Warnln *
func Warnln(v ...interface{}) {
	log.Warnln(v...)
}

//Warnf *
func Warnf(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

//Error *
func Error(v ...interface{}) {
	log.Error(v...)
}

//Errorln *
func Errorln(v ...interface{}) {
	log.Errorln(v...)
}

//Errorf *
func Errorf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

//Fatal *
func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

//Fatalln *
func Fatalln(v ...interface{}) {
	log.Fatalln(v...)
}

//Fatalf *
func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

//Panic *
func Panic(v ...interface{}) {
	log.Panic(v...)
}

//Panicln *
func Panicln(v ...interface{}) {
	log.Panicln(v...)
}

//Panicf *
func Panicf(format string, v ...interface{}) {
	log.Panicf(format, v...)
}
