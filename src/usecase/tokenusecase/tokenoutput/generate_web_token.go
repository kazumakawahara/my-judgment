package tokenoutput

type GenerateWebTokenOutput struct {
	Token string `json:"token"`
}

func NewGenerateWebTokenOutput(token string) *GenerateWebTokenOutput {
	return &GenerateWebTokenOutput{
		Token: token,
	}
}
