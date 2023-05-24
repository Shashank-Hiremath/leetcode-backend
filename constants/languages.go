package constants

const (
	// UNKNOWN    string = "unknown"
	C          string = "c"
	CPP               = "cpp"
	PYTHON            = "python"
	GOLANG            = "golang"
	JAVA              = "java"
	JAVASCRIPT        = "javascript"
)

var LanguagesMap = map[string]string{
	C:          C,
	CPP:        CPP,
	PYTHON:     PYTHON,
	GOLANG:     GOLANG,
	JAVA:       JAVA,
	JAVASCRIPT: JAVASCRIPT,
}
