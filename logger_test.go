package logger

import (
	"testing"
	"fmt"
	"os"
)

func TestStdLog(t *testing.T)  {
	cfg := Config("./logcfg.json")
	Bind(cfg)
	if file,err := os.OpenFile(cfg.Prefix + "log.txt",os.O_APPEND | os.O_CREATE,0766); nil == err {
		SetOutput(file)
	} else {
		fmt.Println(err.Error())
	}
	Debug("std:",cfg.String())
	Info("std dbginfo")
	Warn("std warndbginfo")
	GetLog().Print("std:",123)
	GetLog().WithFields(RunInfo()).Debug("std:yes")
}

func TestNewLog(t *testing.T)  {
	cfg := Config("./logcfg.json")
	l := New("",5,3,"2006-01-02 15:04:05",3)
	l.Bind(cfg)

	if file,err := os.OpenFile(cfg.Prefix + "log.txt",os.O_APPEND | os.O_CREATE,0766); nil == err {
		l.SetOutput(file)
	} else {
		fmt.Println(err.Error())
	}

	l.Debug(cfg.String())
	l.Debug(111,345,"fdsfsdf",555,666)
	l.Info("dbginfo")
	l.Warn("warndbginfo")
	l.GetLog().Print(123)
	l.GetLog().WithFields(l.RunInfo(2)).Debug("yes")
}
