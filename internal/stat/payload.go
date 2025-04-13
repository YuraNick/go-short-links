package stat

// type StatGetRequest struct {
// 	From datatypes.Date `json:"from" validate:"required"`
// 	To   datatypes.Date `json:"to" validate:"required"`
// 	By   string         `json:"by" validate:"required"`
// }

type GetStatResponse struct {
	Period string `json:"period"`
	Sum    int    `json:"sum"`
}
