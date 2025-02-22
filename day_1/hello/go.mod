module example.com/hello

go 1.24.0

replace example.com/greetings => ../greetings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	example.com/hbday v0.0.0-00010101000000-000000000000
)

replace example.com/hbday => ../hbday
