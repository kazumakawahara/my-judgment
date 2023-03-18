package mjerr

type OriginError interface {
	Error() string
	Level() Level
}
