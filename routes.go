package main

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		IndexHandler,
	},
	Route{
		"CreateContent",
		"POST",
		"/content",
		CreateContentReservationHandler,
	},
	Route{
		"ReadContent",
		"GET",
		"/content/{id}",
		ReadContentHandler,
	},
	Route{
		"UpdateContent",
		"PUT",
		"/content/{id}",
		UpdateContentHandler,
	},
	Route{
		"ListTags",
		"GET",
		"/tag",
		ListTagsHandler,
	},
	Route{
		"CreateTag",
		"POST",
		"/app/{app-uuid}/tag/{tag}",
		CreateTagHandler,
	},
	Route{
		"DeleteTag",
		"DELETE",
		"/tag/{id}",
		DeleteTagHandler,
	},
	Route{
		"ReadNewestLocalizedContentEntriesForTag",
		"GET",
		"/app/{app-uuid}/tag/{tag}/locale/{locale}",
		ReadNewestLocalizedContentEntriesForTagHandler,
	},
	Route{
		"ReadAllContentForTag",
		"GET",
		"/app/{app-uuid}/tag/{tag}",
		ReadAllContentForTagHandler,
	},
	Route{
		"AssignTagToContent",
		"POST",
		"/content/{content-uuid}/tag/{tag}",
		AssignTagToContentHandler,
	},
	Route{
		"Login",
		"POST",
		"/auth/login",
		LoginHandler,
	},
	Route{
		"Logout",
		"DELETE",
		"/auth/logout",
		LogoutHandler,
	},
	Route{
		"CreateUser",
		"POST",
		"/auth/user",
		CreateUserHandler,
	},
	Route{
		"GetUser",
		"GET",
		"/auth/user",
		GetUserHandler,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/auth/user",
		DeleteUserHandler,
	},
	Route{
		"AuthenticateToken",
		"POST",
		"/auth/token",
		AuthenticateTokenHandler,
	},
	Route{
		"NewUUID",
		"GET",
		"/auth/uuid",
		NewUUIDHandler,
	},
	Route{
		"AssetUploadURL",
		"POST",
		"/asset/url",
		AssetUploadURLHandler,
	},
	Route{
		"UpdatePassword",
		"PUT",
		"/auth/user/password",
		UpdatePasswordHandler,
	},
	Route{
		"ForgotPassword",
		"POST",
		"/auth/user/password",
		ForgotPasswordHandler,
	},
	Route{
		"UserCryptoBootstrap",
		"POST",
		"/auth/crypto",
		UserCryptoBootstrapHandler,
	},
}
