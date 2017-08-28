package logger

import (
	log "github.com/sirupsen/logrus"
	"fmt"
	"encoding/json"
	"reflect"
	"io/ioutil"
	"runtime"
	"sync"
	"io"
	"os"
)

type Logger struct {
	lock sync.Mutex
	log  *log.Logger
	cfg  LogConfig
	depth int
}

type LogConfig struct {
	//true is json otherwise is text
	Format        bool
	//file prefix
	Prefix        string
	//log level
	Level         log.Level
	//open log flag
	IsOpen        bool
	//log file func line flag
	Flag          int
	//json format
	JsonFormatter log.JSONFormatter
	//text format
	TextFormatter log.TextFormatter
}

var std = New("./DebugLogs/",log.DebugLevel,3,"2006-01-02 15:04:05",4)

func New(prefix string,level log.Level,flag int,TimestampFormat string,depth int) *Logger {
	l := &Logger{
		log : log.New(),
		cfg : LogConfig{
			true,
			prefix,
			level,
			true,
			flag,
			log.JSONFormatter{TimestampFormat:TimestampFormat,DisableTimestamp:false,FieldMap: log.FieldMap{"FieldKeyTime": "@timestamp","FieldKeyLevel": "@level","FieldKeyMsg": "@message"}},
			log.TextFormatter{ForceColors:true,DisableColors:true,DisableTimestamp:false,FullTimestamp:true,TimestampFormat:TimestampFormat,DisableSorting:false,QuoteEmptyFields: true},
		},
		depth:depth,
	}
	set(l.log,l.cfg)
	return l
}

func (l *Logger)Bind(cfg LogConfig)  {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.cfg = cfg
	set(l.log,l.cfg)
}

func Bind(cfg LogConfig)  {
	std.Bind(cfg)
}

func (l *Logger) GetLog() *log.Logger  {
	return l.log
}

func GetLog() *log.Logger  {
	return std.log
}

func (l *Logger)SetOutput(out io.Writer)  {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.log.Out = out
}

func SetOutput(out io.Writer)  {
	std.SetOutput(out)
}

func (l *Logger)SetLevel(level log.Level) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.cfg.IsOpen {
		l.log.Level = level
		l.cfg.Level = level
	} else {
		l.log.Level = 0
		l.cfg.Level = 0
	}
}

func SetLevel(level log.Level)  {
	std.SetLevel(level)
}

func (l *Logger)SetFormatter(formatter log.Formatter) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.log.Formatter = formatter
}

func SetFormatter(formatter log.Formatter) {
	std.SetFormatter(formatter)
}

func (l *Logger)SetOpenFlag(open bool)  {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.cfg.IsOpen = open
}

func SetOpenFlag(open bool)  {
	std.SetOpenFlag(open)
}

func (l *Logger)IsOpenFlag()  bool {
	return l.cfg.IsOpen
}

func IsOpenFlag()  bool {
	return std.IsOpenFlag()
}

func (l *Logger) Flags() int {
	return l.cfg.Flag
}

func Flags() int {
	return std.cfg.Flag
}

func (l *Logger) SetFlags(flag int) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.cfg.Flag = flag
}

func SetFlags(flag int) {
	std.SetFlags(flag)
}

func (l *Logger) Debug(args ...interface{}) {
	l.write(log.DebugLevel,l.withRunInfoFields().Debug,args...)
}

func Debug(args ...interface{})  {
	std.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.write(log.InfoLevel,l.withRunInfoFields().Info,args...)
}

func Info(args ...interface{})  {
	std.Info(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.write(log.ErrorLevel,l.withRunInfoFields().Error,args...)
}

func Error(args ...interface{})  {
	std.Error(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.write(log.WarnLevel,l.withRunInfoFields().Warn,args...)
}

func Warn(args ...interface{})  {
	std.Warn(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.write(log.FatalLevel,l.withRunInfoFields().Fatal,args...)
}

func Fatal(args ...interface{})  {
	std.Fatal(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.write(log.PanicLevel,l.withRunInfoFields().Panic,args)
}

func Panic(args ...interface{}) {
	std.Panic(args)
}

func (l *Logger) RunInfo(depth int) log.Fields  {
	return runInfo(depth,l.cfg.Flag)
}

func RunInfo() log.Fields  {
	return runInfo(2,std.cfg.Flag)
}

func runInfo(depth int,flag int) log.Fields  {
	if flag & (llongfile | lshortfile | lfunc | lline) == 0 {
		return log.Fields{}
	}
	rst := make(map[string]interface{})
	if funcname,file,line,ok := runtime.Caller(depth); ok {
		if flag & lshortfile != 0 {
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					file = file[i + 1:]
					break
				}
			}
		}
		if flag & (llongfile | lshortfile) != 0 {
			rst["file"] = file
		}
		if flag & lline != 0 {
			rst["line"] = line
		}
		if flag & lfunc != 0 {
			funcName := runtime.FuncForPC(funcname).Name()
			for i := len(funcName) - 1; i > 0; i-- {
				if funcName[i] == '/' {
					funcName = funcName[i + 1:]
					break
				}
			}
			rst["func"] = funcName
		}
	}

	return rst
}

func (l LogConfig) String() string {
	rst := ""
	obj := reflect.ValueOf(&l)
	elm := obj.Elem()
	etype := elm.Type()
	rst += fmt.Sprintf("%s { \n",etype.Name())
	for i := 0; i < elm.NumField(); i++ {
		filed := elm.Field(i)
		rst += fmt.Sprintf("	%s %s = %v \n", etype.Field(i).Name,filed.Type(),filed.Interface())
	}
	rst += "}"
	return rst
}

func Config(filename string) LogConfig {
	cfg := LogConfig{
		true,
		"./DebugLogs/",
		log.DebugLevel,
		true,
		lshortfile | lline,
		log.JSONFormatter{TimestampFormat:"2006-01-02 15:04:05",DisableTimestamp:false,FieldMap: log.FieldMap{"FieldKeyTime": "@timestamp","FieldKeyLevel": "@level","FieldKeyMsg": "@message"}},
		log.TextFormatter{ForceColors:true,DisableColors:true,DisableTimestamp:false,FullTimestamp:true,TimestampFormat:"2006-01-02 15:04:05",DisableSorting:false,QuoteEmptyFields: true},
	}

	if ok,_ := exists(filename); ok {
		if buf,err := ioutil.ReadFile(filename); nil != err {
			fmt.Println(err.Error())
		} else {
			if err := json.Unmarshal(buf,&cfg); nil != err {
				fmt.Println(err.Error())
			} else {
				fmt.Println(cfg.String())
			}
		}
	} else {
		fmt.Println(filename,"isn't exists!")
	}

	return cfg
}

func (l *Logger)filter(lvl log.Level) bool  {
	if !l.cfg.IsOpen || l.cfg.Level < lvl {
		return true
	}
	return false
}

func set(l *log.Logger,cfg LogConfig)  {
	if cfg.Format {
		l.Formatter = &cfg.JsonFormatter
	} else {
		l.Formatter = &cfg.TextFormatter
	}
	if cfg.IsOpen {
		l.Level = cfg.Level
	} else {
		l.Level = 0
	}
}

func (l *Logger) write(level log.Level,callback func(args ...interface{}),args ...interface{})  {
	if (l.filter(level)) {
		return
	}
	callback(args...)
}

func (l *Logger) withRunInfoFields() *log.Entry {
	return l.log.WithFields(runInfo(l.depth,l.cfg.Flag))
}

//file or path whether exists
func exists(path string) (bool,error)  {
	if _,err := os.Stat(path); nil == err {
		return true,nil
	} else {
		if os.IsNotExist(err) {
			return false,nil
		} else {
			return false,err
		}
	}
}

const (
	lshortfile = 1 << iota
	lline
	lfunc
	llongfile
)