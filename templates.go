package main

var upload =
`<html>
	<head>
		<title>Upload File</title>
	</head>
	<body>
	<form enctype="multipart/form-data" action="#" method="POST">
		<input type="file" name="file"/>
		<input type="submit" value="upload"/>
	</form>
	</body>
</html> `

var frame = 
`<html>
	<head>
		<title>NetShare</title>
	</head>
	<body>
		<h1>NetShare (<a href="/upload">upload</a>) </h1> 
		<iframe src="/browse" height="100%" width="100%"></iframe>
	</body>
</html> `
