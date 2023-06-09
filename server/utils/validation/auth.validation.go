package validation

import (
	"fmt"
	"regexp"
	"strings"
	"sample-table/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type AuthValidation interface {
	SignUpValidate(requestParams model.SignUpRequest) error
}

type authValidation struct {}

func NewAuthValidation() AuthValidation {
	return &authValidation{}
}

func JPChar(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("%vは文字列でなければなりません", value)
	}

	pattern := "^[ぁ-ゔァ-ヴー々〆〤-\u309f\u30a0-\u30ff\u3400-\u4dbf\u4e00-\u9fff\uf900-\ufaff]*$"
	re := regexp.MustCompile(pattern)

	if !re.MatchString(s) {
		return fmt.Errorf("%vは日本語のみの入力でなければなりません", s)
	}
	return nil
}



//ドメインの確認
func DomainCheck(domain string) validation.RuleFunc {
	return func(value interface{}) error {
		s, ok := value.(string)
		if(!ok){
			return validation.NewError("validation_string", fmt.Sprintf("%vは文字列でなければなりません", value))
		}

		//ドメインが正しいかどうか
		parts := strings.Split(s, "@")
		if len(parts) != 2 || parts[1] != domain {
			return validation.NewError("validation_domain", fmt.Sprintf("メールアドレスのドメインは%vである必要があります。", domain))
		}

		return nil
	}
}

//サインアップ時のバリデーション機能
func (av *authValidation) SignUpValidate(requestParams model.SignUpRequest) error {
	return validation.ValidateStruct(&requestParams,
		validation.Field(
			&requestParams.FirstName,
			validation.Required.Error("お名前は必須入力です。"),
			validation.By(JPChar),
		),
		validation.Field(
			&requestParams.LastName,
			validation.Required.Error("お名前は必須入力です。"),
			validation.By(JPChar),
			),
		validation.Field(
			&requestParams.Email,
			validation.Required.Error("メールアドレスは必須です"),
			validation.By(DomainCheck("gmail.com")),
		),
		validation.Field(
			&requestParams.Password,
			validation.Required.Error("パスワードは必須です。"),
			validation.RuneLength(8, 20).Error("パスワードは8文字以上20文字いないでなけれななりません。"),
			is.Alphanumeric.Error("パスワードは英数字でなければなりません。"),
		),
	)
}
