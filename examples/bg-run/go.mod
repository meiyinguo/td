module github.com/gotd/td/examples/bg-run

go 1.16

require (
	github.com/gotd/contrib v0.15.0
	github.com/gotd/td v0.77.0
	go.uber.org/zap v1.24.0
)

replace github.com/gotd/td => ./../..
