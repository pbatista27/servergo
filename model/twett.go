package model

type Twett struct {
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
}
