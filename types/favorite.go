package types

type FavoritesResponse struct {
	Response
	VideoList []*Video `json:"video_list"`
}
