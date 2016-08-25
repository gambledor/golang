<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>{{.Title}}</title>
</head>
<body>
	<form action="/login" method="post" accept-charset="utf-8">
		Username: <input type="text" name="username">
		Passower: <input type="password" name="password">
		<input type="submit" value="Login">
	</form>
</body>
</html>
