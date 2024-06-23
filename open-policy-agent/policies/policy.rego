package authz

default allow := false

# Allow requests if the user is "admin"
allow if {
	input.user == "admin"
}

# Allow requests if the user has the required permission
allow if {
	user_has_permission(input.user, input.action, input.resource)
}

user_has_permission(user, action, resource) if {
	permissions := data.user_permissions[user]
	resource in data.user_permissions[user][action]
}

