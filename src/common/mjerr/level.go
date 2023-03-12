package mjerr

type Level string

const (
	LevelWarning Level = "WARNING"
	LevelError   Level = "ERROR"
	LevelFatal   Level = "FATAL"
)

func (e *mjError) IsLevelWarning() bool {
	return e.checkLevel(LevelWarning)
}

func (e *mjError) IsLevelError() bool {
	return e.checkLevel(LevelError)
}

func (e *mjError) IsLevelFatal() bool {
	return e.checkLevel(LevelFatal)
}

func (e *mjError) checkLevel(lv Level) bool {
	if e.originError != nil {
		return e.originError.Level() == lv
	}

	if next := AsApoError(e.next); next != nil {
		return next.checkLevel(lv)
	}

	// Default level is fatal
	return lv == LevelFatal
}
