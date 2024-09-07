package web

type KelasCreateRequest struct {
	Ruang string `validate:"required" json:"name"`
}
