package favourite

//Favourite define the relationship between struct user and movie
type Favourite struct {
	IDUser  int `db:"idUser"`
	IDMovie int `db:"idMovie"`
}
