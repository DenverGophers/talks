
# Demo 1

## Automate build, test and lint/vet for go files using guard-shell.

Configure our Gemfile (bundler)
	Add "gem 'guard-shell'" to Gemfile
	run `bundle`

Configure our Guardfile (gaurd)
	teach it to watch go files and react
	run `guard --no-interactions`

Edit go files and watch our console for complaints
