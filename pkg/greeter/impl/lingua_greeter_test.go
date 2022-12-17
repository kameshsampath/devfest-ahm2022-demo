package impl

import (
	"context"
	"github/kameshsampath/devfest-ahm22/pkg/greeter"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestOneLangTrans(t *testing.T) {
	target := "localhost:9090"
	if t, ok := os.LookupEnv("TARGET_SERVICE"); ok {
		target = t
	}
	bgCtx := context.Background()
	ctx, cancel := context.WithTimeout(bgCtx, 10*time.Second)
	defer cancel()

	cc, err := grpc.DialContext(ctx, target, grpc.WithTransportCredentials(insecure.NewCredentials()))

	assert.NoError(t, err)

	c := greeter.NewGreeterClient(cc)
	count := 0
	for {
		if count > 0 {
			break
		}

		res, err := c.Greet(bgCtx, &greeter.GreetRequest{
			Message:     "Hello World!",
			SourceLang:  language.English.String(),
			TargetLangs: []string{language.Gujarati.String()},
		})

		assert.NoError(t, err)

		got, err := res.Recv()
		assert.NoError(t, err)
		assert.Equal(t, "હેલો વર્લ્ડ!", got.Message)
		assert.Equal(t, language.Gujarati.String(), got.Lang)

		count++
	}
}
