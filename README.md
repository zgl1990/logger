# logger
Implement configuration json file Encapsulation Logrus

Logrus had't configure file，I need to config log when program is running，so I make it。

# Test
this dir had a configure file and filename is logcfg.json
configure file content like this below：

{
	"format": true,
	"Prefix":"./",
	"Level":5,
	"IsOpen":true,
	"Flag":3,
	"JsonFormatter":{"TimestampFormat":"2006-01-02 15:04:05","DisableTimestamp":false,"FieldMap":{"FieldKeyTime": "@timestamp","FieldKeyLevel": "@level","FieldKeyMsg": "@message"}},
	"TextFormatter":{"ForceColors":false,"DisableColors":true,"DisableTimestamp":false,"FullTimestamp":true,"TimestampFormat":"2006-01-02 15:04:05","DisableSorting":false,"QuoteEmptyFields": true}
}

run "go test logger" will output a file this dir and filename is log.txt with content:
{"file":"logger_test.go","level":"debug","line":17,"msg":"std:LogConfig { \n\tFormat bool = true \n\tPrefix string = ./ \n\tLevel logrus.Level = debug \n\tIsOpen bool = true \n\tFlag int = 3 \n\tJsonFormatter logrus.JSONFormatter = {2006-01-02 15:04:05 false map[FieldKeyTime:@timestamp FieldKeyLevel:@level FieldKeyMsg:@message]} \n\tTextFormatter logrus.TextFormatter = {false true false true 2006-01-02 15:04:05 false true false {{0 0} 0}} \n}","time":"2017-08-28 09:53:56"}
{"file":"logger_test.go","level":"info","line":18,"msg":"std dbginfo","time":"2017-08-28 09:53:56"}
{"file":"logger_test.go","level":"warning","line":19,"msg":"std warndbginfo","time":"2017-08-28 09:53:56"}
{"level":"info","msg":"std:123","time":"2017-08-28 09:53:56"}
{"file":"logger_test.go","level":"debug","line":21,"msg":"std:yes","time":"2017-08-28 09:53:56"}
{"file":"logger_test.go","level":"debug","line":35,"msg":"LogConfig { \n\tFormat bool = true \n\tPrefix string = ./ \n\tLevel logrus.Level = debug \n\tIsOpen bool = true \n\tFlag int = 3 \n\tJsonFormatter logrus.JSONFormatter = {2006-01-02 15:04:05 false map[FieldKeyTime:@timestamp FieldKeyLevel:@level FieldKeyMsg:@message]} \n\tTextFormatter logrus.TextFormatter = {false true false true 2006-01-02 15:04:05 false true false {{0 0} 0}} \n}","time":"2017-08-28 09:53:56"}
{"file":"logger_test.go","level":"debug","line":36,"msg":"111 345fdsfsdf555 666","time":"2017-08-28 09:53:56"}
{"file":"logger_test.go","level":"info","line":37,"msg":"dbginfo","time":"2017-08-28 09:53:56"}
{"file":"logger_test.go","level":"warning","line":38,"msg":"warndbginfo","time":"2017-08-28 09:53:56"}
{"level":"info","msg":"123","time":"2017-08-28 09:53:56"}
{"file":"logger_test.go","level":"debug","line":40,"msg":"yes","time":"2017-08-28 09:53:56"}

# other
you can use it like Logrus through GetLog() but std，because std is private
