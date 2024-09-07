package web

type KelasCreateResponse struct {
	Id    int    `validate:"required" json:"id"`
	Ruang string `validate:"required" json:"name"`
}
