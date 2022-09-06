package users

type User struct {
	Email    string
	Password string
}
type authUser struct {
	email        string
	passwordHash string
}

var authUserDB = map[string]authUser{}	// email => authUser{email, hash}
var DefaultUserService UserService
type UserService struct {}

func (UserService) verifyUser(user User) bool {
	authUser, ok := authUserDB[user.Email]
	if !ok {
		return false
	}
	err := bcrypt.CompareHashAndPassword(
		[]byte(authUser.passwordHash),
		[]byte(user.Password),
	)
	return err == nil
}

func (UserService) createUser(newUser User) {
	_, ok := authUserDB[newUser.Email]
	if ok {
		fmt.Println("user already exists")
		return errors.new("user already exists")
	}
	passwordHash, err := getPasswordHash(newUser.Password)
	if err != nil {
		return err
	}
	newAuthUser := authUser{
		email: newUser.Email,
		passwordHash: getPasswordHash(newUser.Password),
	}
	authUserDB[newAuthUser.email] = newAuthUser
}

getPasswordHash(password string) (string, error) {
	hash, error := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(hash), err
}

getUser(r *http.Request) users.User {
	email := r.FormValue("email")
	password := r.FormValue("password")
	return users.User{
		Email: email,
		Password: password,
	}
}