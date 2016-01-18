package log

import "io"

// 可以用这些串和日期、时间（包含毫秒数）任意组合，拼成各种格式的日志，如 csv/json/xml
const (
	LevelToken   string = "info"
	TagToken            = "tag"
	PathToken           = "/go/src/github.com/gotips/log/examples/main.go"
	PackageToken        = "github.com/gotips/log/examples/main.go"
	ProjectToken        = "examples/main.go"
	FileToken           = "main.go"
	LineToken    int    = 88
	MessageToken string = "message"
)

// DefaultFormat 默认日志格式
const DefaultFormat = "2006-01-02 15:04:05 info examples/main.go:88 message"

// DefaultFormatTag 默认日志格式带标签
const DefaultFormatTag = "2006-01-02 15:04:05 tag info examples/main.go:88 message"

// Level 日志级别
type Level uint8

// 所有日志级别常量，级别越高，日志越重要，级别越低，日志越详细
const (
	AllLevel Level = iota // 等同于 TraceLevel

	TraceLevel
	DebugLevel // 默认日志级别，方便开发
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel // panic 打印错误栈，但是可以 recover
	FatalLevel // fatal 表明严重错误，程序直接退出，慎用

	// 提供这个级别日志，方便在无论何种情况下，都打印必要信息，比如服务启动信息
	PrintLevel
	StackLevel // 打印堆栈信息
)

// Labels 每个级别对应的标签
var Labels = [...]string{"all", "trace", "debug", "info", "warn", "error", "panic", "fatal", "print", "stack"}

// String 返回日志标签
func (v Level) String() string {
	return Labels[v]
}

var (
	// 默认 debug 级别，方便调试，生产环境可以调用 SetLevel 设置 log 级别
	v Level = DebugLevel
	// 默认实现，输出到 os.Std 中，可以重定向到文件中，也可以调用 SetPrinter 其他方式输出
	std Printer
)

// SetLevel 设置日志级别
func SetLevel(l Level) { v = l }

// SetPrinter 切换 Printer 实现
func SetPrinter(p Printer) { std = p }

// ChangeWriter 改变输出位置，通过这个接口，可以实现日志文件按时间或按大小滚动
func ChangeWriter(w io.Writer) { std.ChangeWriter(w) }

// ChangeFormat 改变日志格式
func ChangeFormat(format string) { std.ChangeFormat(format) }

// 判断各种级别的日志是否会被输出
func IsTraceEnabled() bool { return v <= TraceLevel }
func IsDebugEnabled() bool { return v <= DebugLevel }
func IsInfoEnabled() bool  { return v <= InfoLevel }
func IsWarnEnabled() bool  { return v <= WarnLevel }
func IsErrorEnabled() bool { return v <= ErrorLevel }
func IsPanicEnabled() bool { return v <= PanicLevel }
func IsFatalEnabled() bool { return v <= FatalLevel }
func IsPrintEnabled() bool { return v <= PrintLevel }
func IsStackEnabled() bool { return v <= StackLevel }

// 打印日志
func Trace(m ...interface{}) { std.Tprintf(v, TraceLevel, "", "", m...) }
func Debug(m ...interface{}) { std.Tprintf(v, DebugLevel, "", "", m...) }
func Info(m ...interface{})  { std.Tprintf(v, InfoLevel, "", "", m...) }
func Warn(m ...interface{})  { std.Tprintf(v, WarnLevel, "", "", m...) }
func Error(m ...interface{}) { std.Tprintf(v, ErrorLevel, "", "", m...) }
func Panic(m ...interface{}) { std.Tprintf(v, PanicLevel, "", "", m...) }
func Fatal(m ...interface{}) { std.Tprintf(v, FatalLevel, "", "", m...) }
func Print(m ...interface{}) { std.Tprintf(v, PrintLevel, "", "", m...) }
func Stack(m ...interface{}) { std.Tprintf(v, StackLevel, "", "", m...) }

// 按一定格式打印日志
func Tracef(format string, m ...interface{}) { std.Tprintf(v, TraceLevel, "", format, m...) }
func Debugf(format string, m ...interface{}) { std.Tprintf(v, DebugLevel, "", format, m...) }
func Infof(format string, m ...interface{})  { std.Tprintf(v, InfoLevel, "", format, m...) }
func Warnf(format string, m ...interface{})  { std.Tprintf(v, WarnLevel, "", format, m...) }
func Errorf(format string, m ...interface{}) { std.Tprintf(v, ErrorLevel, "", format, m...) }
func Panicf(format string, m ...interface{}) { std.Tprintf(v, PanicLevel, "", format, m...) }
func Fatalf(format string, m ...interface{}) { std.Tprintf(v, FatalLevel, "", format, m...) }
func Printf(format string, m ...interface{}) { std.Tprintf(v, PrintLevel, "", format, m...) }
func Stackf(format string, m ...interface{}) { std.Tprintf(v, StackLevel, "", format, m...) }

// 打印日志时带上 tag
func Ttrace(tag string, m ...interface{}) { std.Tprintf(v, TraceLevel, tag, "", m...) }
func Tdebug(tag string, m ...interface{}) { std.Tprintf(v, DebugLevel, tag, "", m...) }
func Tinfo(tag string, m ...interface{})  { std.Tprintf(v, InfoLevel, tag, "", m...) }
func Twarn(tag string, m ...interface{})  { std.Tprintf(v, WarnLevel, tag, "", m...) }
func Terror(tag string, m ...interface{}) { std.Tprintf(v, ErrorLevel, tag, "", m...) }
func Tpanic(tag string, m ...interface{}) { std.Tprintf(v, PanicLevel, tag, "", m...) }
func Tfatal(tag string, m ...interface{}) { std.Tprintf(v, FatalLevel, tag, "", m...) }
func Tprint(tag string, m ...interface{}) { std.Tprintf(v, PrintLevel, tag, "", m...) }
func Tstack(tag string, m ...interface{}) { std.Tprintf(v, StackLevel, tag, "", m...) }

// 按一定格式打印日志，并在打印日志时带上 tag
func Ttracef(tag string, format string, m ...interface{}) {
	std.Tprintf(v, TraceLevel, tag, format, m...)
}
func Tdebugf(tag string, format string, m ...interface{}) {
	std.Tprintf(v, DebugLevel, tag, format, m...)
}
func Tinfof(tag string, format string, m ...interface{}) { std.Tprintf(v, InfoLevel, tag, format, m...) }
func Twarnf(tag string, format string, m ...interface{}) { std.Tprintf(v, WarnLevel, tag, format, m...) }
func Terrorf(tag string, format string, m ...interface{}) {
	std.Tprintf(v, ErrorLevel, tag, format, m...)
}
func Tpanicf(tag string, format string, m ...interface{}) {
	std.Tprintf(v, PanicLevel, tag, format, m...)
}
func Tfatalf(tag string, format string, m ...interface{}) {
	std.Tprintf(v, FatalLevel, tag, format, m...)
}
func Tprintf(tag string, format string, m ...interface{}) {
	std.Tprintf(v, PrintLevel, tag, format, m...)
}
func Tstackf(tag string, format string, m ...interface{}) {
	std.Tprintf(v, StackLevel, tag, format, m...)
}

// ======== 兼容 wothing/log ===============

// 打印日志时带上 tag
func TraceT(tag string, m ...interface{}) { std.Tprintf(v, TraceLevel, tag, "", m...) }
func DebugT(tag string, m ...interface{}) { std.Tprintf(v, DebugLevel, tag, "", m...) }
func InfoT(tag string, m ...interface{})  { std.Tprintf(v, InfoLevel, tag, "", m...) }
func WarnT(tag string, m ...interface{})  { std.Tprintf(v, WarnLevel, tag, "", m...) }
func ErrorT(tag string, m ...interface{}) { std.Tprintf(v, ErrorLevel, tag, "", m...) }
func PanicT(tag string, m ...interface{}) { std.Tprintf(v, PanicLevel, tag, "", m...) }
func FatalT(tag string, m ...interface{}) { std.Tprintf(v, FatalLevel, tag, "", m...) }
func PrintT(tag string, m ...interface{}) { std.Tprintf(v, PrintLevel, tag, "", m...) }
func StackT(tag string, m ...interface{}) { std.Tprintf(v, StackLevel, tag, "", m...) }

// 按一定格式打印日志，并在打印日志时带上 tag
func TracefT(tag string, format string, m ...interface{}) {
	std.Tprintf(v, TraceLevel, tag, format, m...)
}
func DebugfT(tag string, format string, m ...interface{}) {
	std.Tprintf(v, DebugLevel, tag, format, m...)
}
func InfofT(tag string, format string, m ...interface{}) { std.Tprintf(v, InfoLevel, tag, format, m...) }
func WarnfT(tag string, format string, m ...interface{}) { std.Tprintf(v, WarnLevel, tag, format, m...) }
func ErrorfT(tag string, format string, m ...interface{}) {
	std.Tprintf(v, ErrorLevel, tag, format, m...)
}
func PanicfT(tag string, format string, m ...interface{}) {
	std.Tprintf(v, PanicLevel, tag, format, m...)
}
func FatalfT(tag string, format string, m ...interface{}) {
	std.Tprintf(v, FatalLevel, tag, format, m...)
}
func PrintfT(tag string, format string, m ...interface{}) {
	std.Tprintf(v, PrintLevel, tag, format, m...)
}
func StackfT(tag string, format string, m ...interface{}) {
	std.Tprintf(v, StackLevel, tag, format, m...)
}
