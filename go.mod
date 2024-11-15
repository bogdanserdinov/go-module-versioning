module github.com/bogdanserdinov/sdk-versioning-test

go 1.22.8

// Retract version v1.0.0 due to breaking changes
retract v1.0.0

require github.com/stretchr/testify v1.9.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
