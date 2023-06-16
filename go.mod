module github.com/liang09255/lutils

go 1.20

require github.com/stretchr/testify v1.8.4

replace (
	github.com/liang09255/lutils/conv => ./conv
	github.com/liang09255/lutils/dfo => ./dfo
)

require (
	github.com/liang09255/lutils/conv v1.0.0
	github.com/liang09255/lutils/dfo v1.0.0
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

