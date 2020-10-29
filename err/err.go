package err

func Err(s string) error {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Errorf("%s: %s:%d %s\n", s, frame.File, frame.Line, frame.Function)
}