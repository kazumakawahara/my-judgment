package userinput

type UpdateUserInput struct {
	UserID int        `param:"userID"`
	User   UpdateUser `json:"mjUser"`
}

type UpdateUser struct {
	Name     *string `json:"name"`
	Gender   *string `json:"gender"`
	Address  *string `json:"address"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}
