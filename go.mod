module github.com/kendfss/pipe

go 1.18

require github.com/kendfss/but v0.0.0-00010101000000-000000000000

replace (
	github.com/kendfss/but => ../but
	github.com/kendfss/iters => ../iters
	github.com/kendfss/oprs => ../oprs
	github.com/kendfss/oracle => ../oracle
	github.com/kendfss/rules => ../rules
)
