module learn-go

go 1.24.0

replace example.com/hello => ./hello

replace example.com/hbday => ./hbday

replace example.com/greetings => ./greetings

replace example.com/greetings-errors => ./greetings-errors

replace example.com/greetings-multiple => ./greetings-multiple

replace greetings-random => ./greetings-random

require example.com/greetings-multiple v0.0.0-00010101000000-000000000000

require greetings-random v0.0.0-00010101000000-000000000000 // indirect
