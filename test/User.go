package testwww

type UserModel struct {
	Name, Sex   string
	PhoneNumber Login
	test        []string
}

// Embed
type Login struct {
	Name string
	Type string
}

type UserT struct {
	Name  string
	Login Login
}

type GirlModel struct {
	Name string
	Time string
}
