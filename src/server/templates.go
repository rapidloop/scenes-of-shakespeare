package main

const tplStrHome = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Scenes of Shakespeare</title>
    <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
	<link href="https://fonts.googleapis.com/css?family=Allura" rel="stylesheet">
  </head>
  <body>
    <div class="container">
	<div class="row" style="text-align: center">
		<h1 style="padding-top: 2em; font-family: Allura, cursive; font-size: 80px">Scenes of Shakespeare</h1>
	</div>
	<div class="row" style="padding-top: 1.5em"><div class="col-sm-4 col-sm-offset-4">
		<form action="/" method="GET"><input class="form-control" autofocus name="q" maxlength=100 type="text"></form>
	</div></div>
	<div class="row" style="text-align: center; padding-top: 3em"><div class="col-sm-6 col-sm-offset-3">
		<p>
		Read this first: <a href="https://www.opsdash.com/blog/postgres-full-text-search-golang.html">Using
		PostgreSQL Full Text Search with Golang</a>. Fork this and use it to build your own full text search tool!
		<p>
		Follow us on twitter <a href="https://twitter.com/therapidloop">@therapidloop</a>.
	</div></div>
    </div>
  </body>
</html>
`

const tplStrResults = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Scenes of Shakespeare</title>
    <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
	<link href="https://fonts.googleapis.com/css?family=Allura" rel="stylesheet">
  </head>
  <body>
    <div class="container-fluid">
	<div class="row" style="text-align: center; background-color: #f1f1f1; padding: 1em 0">
		<div class="col-sm-1" style="font-family: Allura, cursive; font-size: 30px">
		<a href="/" style="text-decoration: none; color: #000">SofS</a>
		</div>
		<div class="col-sm-4">
			<form action="/" method="GET"><input class="form-control" autofocus name="q" maxlength=100 type="text" value="{{.Query}}"></form>
		</div>
	</div>
	{{range .Results}}
	<div class="row" style="padding: 1.5em 0; border-bottom: 1px solid #eee"><div class="col-sm-11 col-sm-offset-1">
		<div style="font-size: 1.2em">
			<a href="/scene?w={{.WorkID}}&a={{.Act}}&s={{.Scene}}">{{.Work}} - Act {{.Act}} Scene {{.Scene}} - {{.Description}}</a>
		</div>
		{{.Snippet}}
	</div></div>
	{{end}}
	<div class="row" style="padding: 1.5em; background-color: #f1f1f1">
		Text and database from <a href="http://opensourceshakespeare.org/">http://opensourceshakespeare.org/</a>
    </div>
    </div>
  </body>
</html>
`

const tplStrScene = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{.Work}} - Act {{.Act}} Scene {{.Scene}}</title>
    <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
	<link href="https://fonts.googleapis.com/css?family=Allura" rel="stylesheet">
  </head>
  <body>
    <div class="container-fluid">
	<div class="row" style="text-align: center; background-color: #f1f1f1; padding: 0.6em 0; font-family: Allura, cursive; font-size: 42px">
		{{.Work}}
	</div>
	<div class="row" style="text-align: center; padding: 1em 0 0 0; font-family: Allura, cursive; font-size: 24px">
		Act {{.Act}} &nbsp; - &nbsp; Scene {{.Scene}}
	</div>
	<div class="row" style="text-align: center; padding: 1em 0; font-family: Allura, cursive; font-size: 28px">
		{{.Description}}
	</div>
	<div class="row"><div class="col-sm-4 col-sm-offset-4">
		{{.Body}}
    </div></div>
	<div class="row" style="padding: 1.5em; background-color: #f1f1f1; text-align: center">
		Text and database from <a href="http://opensourceshakespeare.org/">http://opensourceshakespeare.org/</a>
    </div>
    </div>
  </body>
</html>
`
