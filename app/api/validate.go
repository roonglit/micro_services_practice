package api

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// func validate(user *LoginReq) error {

// validate.RegisterValidation("Email", func(fl validator.FieldLevel) bool {
// 	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
// 	return re.MatchString(fl.Field().String())
// })

// if user.Verified == nil {
// 	user.Verified = new(bool)
// 	*user.Verified = false
// }
// if user.Active == nil {
// 	user.Active = new(bool)
// 	*user.Active = true
// }
// if err := validate.Struct(user); err != nil {
// 	var errMessages []string
// 	for _, err := range err.(validator.ValidationErrors) {
// 		errMessages = append(errMessages, fmt.Sprintf("field %s failed validation: %s", err.Field(), err.Tag()))
// 	}
// 	return errors.New(strings.Join(errMessages, ", "))
// }
// return nil
// }
