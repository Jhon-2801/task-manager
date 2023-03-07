package bd

func CheckExistUser(mail string) (bool, string) {

	user := GetUser(mail)

	if user.Mail == mail {
		return true, user.Password
	}

	return false, user.Password
}
