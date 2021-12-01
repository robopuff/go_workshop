package internal

type App interface {
	Run() error
	ExitCode() int
}