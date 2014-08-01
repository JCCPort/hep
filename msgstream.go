package fwk

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type msgstream struct {
	lvl Level
	w   io.Writer
	n   string
}

func NewMsgStream(name string, lvl Level, w io.Writer) msgstream {
	if w == nil {
		w = os.Stdout
	}
	w = bufio.NewWriter(w)
	return msgstream{
		lvl: lvl,
		w:   w,
		n:   fmt.Sprintf("%-20s", name),
	}
}

func (msg msgstream) Debugf(format string, a ...interface{}) (int, Error) {
	return msg.Msg(LvlDebug, format, a...)
}

func (msg msgstream) Infof(format string, a ...interface{}) (int, Error) {
	return msg.Msg(LvlInfo, format, a...)
}

func (msg msgstream) Warnf(format string, a ...interface{}) (int, Error) {
	return msg.Msg(LvlWarning, format, a...)
}

func (msg msgstream) Errorf(format string, a ...interface{}) (int, Error) {
	return msg.Msg(LvlError, format, a...)
}

func (msg msgstream) Msg(lvl Level, format string, a ...interface{}) (int, Error) {
	if lvl < msg.lvl {
		return 0, nil
	}
	return fmt.Fprintf(msg.w, msg.n+": "+format, a...)
}

// EOF
