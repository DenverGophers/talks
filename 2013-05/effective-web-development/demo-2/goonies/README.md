
# Demo 2

## Automate start/restart your application using guard-process.

Configure our Gemfile (bundler)
	Add "gem 'guard-process'" to Gemfile
	run `bundle`

Configure our Guardfile (gaurd)
	teach it to watch our compiled application and react
	start guard: `guard --no-interactions`

Edit go files and watch our console for complaints

Bonus question
	Why doesn't my count go up sequentially from Chrome?
