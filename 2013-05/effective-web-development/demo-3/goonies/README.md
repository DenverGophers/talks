
# Demo 3

## Automate build for scss files using guard-shell and compass.

Configure our Gemfile (bundler)
	Add "gem 'compass'" to Gemfile
	run `bundle`

Configure our Guardfile (gaurd)
	teach it to watch scss files and react
	run `guard --no-interactions`

Create a directory structure to
	house our scss sources - assets/scss
	serve our compiled css - assets/css (and update our .gitignore)

Update our application
	serve our css file
	render html that includes our css

Edit scss files and watch our console for complaints
