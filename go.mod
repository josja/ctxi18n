module github.com/invopop/ctxi18n

replace github.com/invopop/ctxi18n => github.com/josja/ctxi18n v0.0.0-20251014212415-7cb8ae42feb6
// replace github.com/invopop/ctxi18n => ./

go 1.21

toolchain go1.22.1

require (
	github.com/invopop/yaml v0.2.0
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
