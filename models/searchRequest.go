package models

/*SearchRequest modelo del request*/
type SearchRequest struct {
	Search string `bson:"search" json:"search,omitempty"`
}
