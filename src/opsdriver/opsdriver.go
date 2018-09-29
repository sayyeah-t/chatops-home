package opsdriver

type DriverInterface interface {
	DoCommand([]string) string
}
