package mgembedtest


type Login struct { 
	Name string `bson:"name" json:"name,omitempty"`
	Type string `bson:"type" json:"type,omitempty"`
}

type UserT struct { 
	Name string `bson:"name" json:"name,omitempty"`
	Login Login `bson:"login" json:"login,omitempty"`
}
