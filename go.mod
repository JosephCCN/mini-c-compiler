module main

go 1.23.1

replace github.com/lexical => ./lexical

replace github.com/parser => ./parser

replace github.com/utils => ./utils

require (
	github.com/lexical v0.0.0
	github.com/parser v0.0.0-00010101000000-000000000000
)

require github.com/utils v0.0.0 // indirect
