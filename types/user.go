package types

type User struct {
	ID        string `bson:"_id" json:"id,omitempty"` // omit always json:"_"
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}
