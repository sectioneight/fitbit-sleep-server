package main

import "html/template"

var sleepTemplate = template.Must(template.New("main").Parse(`
<head>
<title>the sleep of ai</title>
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">
</head>
<body>
	<nav class="navbar navbar-inverse navbar-fixed-top">
	  <div class="container">
		</div>
	</nav>

	<div class="container">
		<h1>How much sleep did Aiden get?</h1>
		<p class="lead">{{ .MostRecent }}</p>

		<h5>Date</h5>
		<p>{{ .Date }}</p>

	</div>
</body>
`))
