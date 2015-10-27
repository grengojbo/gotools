package gotools

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAuth(t *testing.T) {
	Convey("Testing encode, decode password\n", t, func() {
		passwordString := GetRandomString(8)
		passwordNum := GetRandomNumString(10)
		Convey("Test bcrypt password", func() {
			pc := PasswordBcrypt(passwordString)
			err := VerifyPassword(pc, passwordString)
			So(err, ShouldBeNil)
			err = VerifyPassword(pc, passwordNum)
			So(err, ShouldNotBeNil)
		})
	})
}
