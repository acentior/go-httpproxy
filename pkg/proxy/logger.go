package proxy

type Logger interface {
	Printf(format string, v ...interface{})
}
