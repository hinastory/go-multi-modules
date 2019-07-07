module go-multi-modules

go 1.12

require (
	github.com/common-nighthawk/go-figure v0.0.0-20190529165535-67e0ed34491a
	go-multi-modules/pkg/go-home v0.0.0-00010101000000-000000000000
)

replace go-multi-modules/pkg/go-home => ./pkg/go-home //追加
