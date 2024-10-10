package Auth

import (
	"fmt"
	"math/rand"
	"net/mail"
	"net/smtp"
	"regexp"
	"strconv"

	"encoding/base64"

	c "ProductService/Connection"
	h "ProductService/Helpers"

	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

func Signin(_email, _password string) (bool, string) {
	var result string
	if IsEmpty(_email) {
		result = `{"statu":"error","message":"Email alanını doldurmanız gerekli."}`
		return h.ExtractStatuAndMessage(result)
	} else if !ValidEmail(_email) {
		result = `{"statu":"error","message":"Geçersiz email formatı."}`
		return h.ExtractStatuAndMessage(result)
	}
	if IsEmpty(_password) {
		result = `{"statu":"error","message":"Şifre alanını doldurmanız gerekli."}`
		return h.ExtractStatuAndMessage(result)
	}
	c.Connection().QueryRow("SELECT signin($1,$2);", _email, base64.StdEncoding.EncodeToString([]byte(_password))).Scan(&result)
	return h.ExtractStatuAndMessage(result)
}
func Signup(_email, _name, _lastname, _nickname, _password, _tel string) (bool, string) {
	var result string
	re, _ := regexp.Compile(`^(\+90\d{10})$`)
	if IsEmpty(_email) {
		result = `{"statu":"error","message":"Email alanını doldurmanız gerekli."}`
		return h.ExtractStatuAndMessage(result)
	} else if !ValidEmail(_email) {
		result = `{"statu":"error","message":"Geçersiz email formatı."}`
		return h.ExtractStatuAndMessage(result)
	}
	if IsEmpty(_name) {
		result = `{"statu":"error","message":"Name alanını doldurmanız gerekli."}`
		return h.ExtractStatuAndMessage(result)
	}
	if IsEmpty(_lastname) {
		result = `{"statu":"error","message":"Lastname alanını doldurmanız gerekli."}`
		return h.ExtractStatuAndMessage(result)
	}
	if IsEmpty(_nickname) {
		result = `{"statu":"error","message":"Nickname alanını doldurmanız gerekli."}`
		return h.ExtractStatuAndMessage(result)
	}
	if IsEmpty(_password) {
		result = `{"statu":"error","message":"Şifre alanını doldurmanız gerekli."}`
		return h.ExtractStatuAndMessage(result)
	}
	if IsEmpty(_tel) {
		result = `{"statu":"error","message":"Telefon numarası alanı doldurulmalı."}`
		return h.ExtractStatuAndMessage(result)
	} else if !re.MatchString(_tel) {
		result = `{"statu":"error","message":"Geçersiz telefon numarası formatı."}`
		return h.ExtractStatuAndMessage(result)
	}
	c.Connection().QueryRow("SELECT signup($1,$2,$3,$4,$5,$6);", _email, _name, _lastname, _nickname, base64.StdEncoding.EncodeToString([]byte(_password)), _tel).Scan(&result)
	return h.ExtractStatuAndMessage(result)
}
func IsAccountVerified(_tel, _email string) (bool, string) {
	var result string
	re, _ := regexp.Compile(`^(\+90\d{10})$`)
	if IsEmpty(_email) {
		result = `{"statu":"error","message":"Email alanını doldurmanız gerekli."}`
		return h.ExtractStatuAndMessage(result)
	} else if !ValidEmail(_email) {
		result = `{"statu":"error","message":"Geçersiz email formatı."}`
		return h.ExtractStatuAndMessage(result)
	}
	if IsEmpty(_tel) {
		result = `{"statu":"error","message":"Telefon numarası alanı doldurulmalı."}`
		return h.ExtractStatuAndMessage(result)
	} else if !re.MatchString(_tel) {
		result = `{"statu":"error","message":"Geçersiz telefon numarası formatı."}`
		return h.ExtractStatuAndMessage(result)
	}
	c.Connection().QueryRow("SELECT isaccountverified($1,$2);", _email, _tel).Scan(&result)
	return h.ExtractStatuAndMessage(result)
}
func Verifications(_email, _tel string, verifytype int) (bool, string) {
	var result string
	re, _ := regexp.Compile(`^(\+90\d{10})$`)
	if IsEmpty(_email) {
		result = `{"statu":"error","message":"Email alanını doldurmanız gerekli."}`
		return h.ExtractStatuAndMessage(result)
	} else if !ValidEmail(_email) {
		result = `{"statu":"error","message":"Geçersiz email formatı."}`
		return h.ExtractStatuAndMessage(result)
	}
	if IsEmpty(_tel) {
		result = `{"statu":"error","message":"Telefon numarası alanı doldurulmalı."}`
		return h.ExtractStatuAndMessage(result)
	} else if !re.MatchString(_tel) {
		result = `{"statu":"error","message":"Geçersiz telefon numarası formatı."}`
		return h.ExtractStatuAndMessage(result)
	}
	if IsEmpty(string(verifytype)) {
		result = `{"statu":"error","message":"Verification alanı doldurulmalı."}`
		return h.ExtractStatuAndMessage(result)
	} else if (verifytype > 2) || (verifytype < 1) {
		result = `{"statu":"error","message":"Geçersiz verification formatı."}`
		return h.ExtractStatuAndMessage(result)
	}

	switch verifytype {
	case 1:
		return SendCodeWithTel(_tel) //sms gönderilecek telefon numarası
	case 2:
		return SendCodeWithEmail(_email)
	}
	return false, "Girilen bilgiler hatalı."
}
func SendCodeWithTel(_tel string) (bool, string) {
	var result string
	re, _ := regexp.Compile(`^(\+90\d{10})$`)
	if IsEmpty(_tel) {
		result = `{"statu":"error","message":"Telefon numarası alanı doldurulmalı."}`
		return h.ExtractStatuAndMessage(result)
	} else if !re.MatchString(_tel) {
		result = `{"statu":"error","message":"Geçersiz telefon numarası formatı."}`
		return h.ExtractStatuAndMessage(result)
	}
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.ConfigModel.Verifmodel.AccountSid,
		Password: c.ConfigModel.Verifmodel.AuthToken,
	})
	var code string = GeneratedCode()
	params := &verify.CreateVerificationParams{}
	params.SetCustomCode(code)
	params.SetChannel("sms")
	params.SetTo(_tel) //Premium olmadıgından tek kişiye sms gönderir
	_, err := client.VerifyV2.CreateVerification(c.ConfigModel.Verifmodel.VerificationString,
		params)
	if err != nil {
		return false, "Telefon kodu gönderilirken hata ile karşılaşıldı.Daha sonra bir daha deneyiniz."
	} else {
		c.Connection().QueryRow("CALL setcode($1,$2);", code, _tel)
		return true, "Telefon kodu başarıyla gönderildi."
	}
}
func SendCodeWithEmail(_email string) (bool, string) {
	var result string
	if IsEmpty(_email) {
		result = `{"statu":"error","message":"Email alanını doldurmanız gerekli."}`
		return h.ExtractStatuAndMessage(result)
	} else if !ValidEmail(_email) {
		result = `{"statu":"error","message":"Geçersiz email formatı."}`
		return h.ExtractStatuAndMessage(result)
	}
	var code string = GeneratedCode()
	auth := smtp.PlainAuth("", c.ConfigModel.Smtp.Address, c.ConfigModel.Smtp.Password, "smtp.gmail.com")
	mesg := fmt.Sprintf("your verification code is %s", code)
	err := smtp.SendMail("smtp.gmail.com:587", auth, c.ConfigModel.Smtp.Address, []string{_email}, []byte(mesg))
	if err != nil {
		return false, "Email gönderilirken hata ile karşılaşıldı.Daha sonra bir daha deneyiniz."
	} else {
		c.Connection().QueryRow("CALL setcode($1,$2);", code, _email)
		return true, "Email başarıyla gönderildi."
	}
}
func VerifyCode(_code string) (bool, string) {
	var result string
	re, _ := regexp.Compile(`^\d{6}$`)
	if IsEmpty(_code) {
		result = `{"statu":"error","message":"Verification kod alanını doldurmanız gerekli."}`
		return h.ExtractStatuAndMessage(result)
	} else if !re.MatchString(_code) {
		result = `{"statu":"error","message":"Verification kod uygun formatta olmalı."}`
		return h.ExtractStatuAndMessage(result)
	}
	c.Connection().QueryRow("SELECT verifyaccount($1);", _code).Scan(&result)
	return h.ExtractStatuAndMessage(result)
}
func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func GeneratedCode() string {
	return strconv.Itoa(rand.Intn(1000000-100000) + 100000)
}
func IsEmpty(value string) bool {
	return len(value) == 0 || value == ""
}
