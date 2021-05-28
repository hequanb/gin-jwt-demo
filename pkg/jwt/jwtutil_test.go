package jwt

import (
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	tokenStr, _ := GenToken("hequan")
	fmt.Println(tokenStr)
}

func TestParseToken(t *testing.T) {
	// tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImhlcXVhbiIsImV4cCI6MTYyMjE4Mzk5NiwiaWF0IjoxNjIyMTc2Nzk2LCJpc3MiOiJnaW4tand0LWRlbW8ifQ.ZwB-H8XWw1v9x7V60cw6JPzyRUDndigwxKlAESr124s"
	fakeTokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImhlcXVIsImV4cCI6MTYyMjE4Mzk5NiwiaWF0IjoxNjIyMTc2Nzk2LCJpc3MiOiJnaW4tand0LWRlbW8ifQ.ZwB-H8XWw1v9x7V60cw6JPzyRUDndigwxKlAESr124s"
	c, err := ParseToken(fakeTokenStr)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", *c)
}
