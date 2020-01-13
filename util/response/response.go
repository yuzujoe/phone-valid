package response

type Response struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type UserAuthenticationSuccessReponse struct {
	Code  int64  `json:"code"`
	Token string `json:"token"`
}
