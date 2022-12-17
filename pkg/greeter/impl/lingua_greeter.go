package impl

import (
	"context"
	"github/kameshsampath/devfest-ahm22/pkg/greeter"
	"log"
	"math/rand"
	"os"
	"time"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

type LinguaGreeterServer struct {
	greeter.UnimplementedGreeterServer
}

// Greet implements greeter.GreeterServer
func (*LinguaGreeterServer) Greet(req *greeter.GreetRequest, stream greeter.Greeter_GreetServer) error {
	apiKey := os.Getenv("DEMO_API_KEY")
	bgCtx := context.Background()
	ctx, cancel := context.WithTimeout(bgCtx, 10*time.Second)
	defer cancel()
	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	for {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(req.TargetLangs), func(i, j int) {
			req.TargetLangs[i], req.TargetLangs[j] = req.TargetLangs[j], req.TargetLangs[i]
		})
		tl := req.TargetLangs[0]
		t, err := client.Translate(bgCtx,
			[]string{req.Message},
			language.MustParse(tl), &translate.Options{
				Source: language.MustParse(req.SourceLang),
				Format: translate.Text,
			})
		if err != nil {
			log.Printf("Error translating %#v,%s", req, err)
		}

		if len(t) > 0 {
			stream.Send(&greeter.GreetResponse{
				Message: t[0].Text,
				Lang:    tl,
			})
			time.Sleep(3 * time.Second)
		}
	}
}

var _ greeter.GreeterServer = (*LinguaGreeterServer)(nil)
