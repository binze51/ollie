package pkg

import (
	"fmt"
	"testing"
)

func TestFeishu(t *testing.T) {
	// info, err := parseToken("21yIdXRPDm67Gni9YL3o/VVLoaXC76yUF9vmavmg8x6893kpqXr2h+uJaqeH2gdaSbWgslJvauNecUxYZLM2iVHQEIIT0bg5iFwVqE0PnIFiM4ol3LWeUgwiHj1DSp1FYoNiE5KQXgvqXLmBII9xczIvNwJe7rRSMLchYNKmYjNMuCo9bijtirmpS3I39d6otR/J4uDQ89YjNq3/OHyD7aUAma8rYNo/a/6wK/1U2++aPqxRSw==")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("%+v", info)

	jwtClaims := &JwtClaims{
		UserName: "heming",
		Phone:    "info.Mobile",
	}
	jwtPayload, _ := CreateJwtToken(jwtClaims)

	fmt.Printf("%+v", jwtPayload)
}
