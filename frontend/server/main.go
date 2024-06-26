package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/curioswitch/go-curiostack/server"
	docshandler "github.com/curioswitch/go-docs-handler"
	protodocs "github.com/curioswitch/go-docs-handler/plugins/proto"
	"github.com/curioswitch/go-usegcp/middleware/firebaseauth"

	frontendapi "github.com/curioswitch/tasuke/frontend/api/go"
	"github.com/curioswitch/tasuke/frontend/api/go/frontendapiconnect"
	"github.com/curioswitch/tasuke/frontend/server/internal/config"
	"github.com/curioswitch/tasuke/frontend/server/internal/handler/saveuser"
	"github.com/curioswitch/tasuke/frontend/server/internal/service"
)

// e2e-test1@curioswitch.org
const e2eTest1UID = "V8yRsCpZJkUfPmxcLI6pKTrx3kf1"

func main() {
	ctx := context.Background()

	conf := config.Load()

	r := server.NewRouter()

	docs, err := docshandler.New(
		protodocs.NewPlugin(
			frontendapiconnect.FrontendServiceName,
			protodocs.WithExampleRequests(
				frontendapiconnect.FrontendServiceSaveUserProcedure,
				&frontendapi.SaveUserRequest{
					User: &frontendapi.User{
						ProgrammingLanguageIds: []uint32{
							132, // golang
						},
						MaxOpenReviews: 5,
					},
				},
			),
		),
		docshandler.WithInjectedScriptSupplier(func() string {
			script := (`
			function include(url) {
				return new Promise((resolve, reject) => {
					var script = document.createElement('script');
					script.type = 'text/javascript';
					script.src = url;

					script.onload = function() {
						resolve({ script });
					};

					document.getElementsByTagName('head')[0].appendChild(script);
				});
			}

			async function loadScripts() {
				await include("https://alpha.tasuke.dev/__/firebase/8.10.1/firebase-app.js");
				await include("https://alpha.tasuke.dev/__/firebase/8.10.1/firebase-auth.js");
				await include("https://alpha.tasuke.dev/__/firebase/init.js");
				firebase.auth();
			}
			loadScripts();

			async function getAuthorization() {
				const token = await firebase.auth().currentUser.getIdToken();
				return {"Authorization": "Bearer " + token};
			}
			window.armeria.registerHeaderProvider(getAuthorization);
			`)
			return script
		}))
	if err != nil {
		log.Fatalf("Failed to create docs handler: %v", err)
	}
	r.Handle("/internal/docs/*", http.StripPrefix("/internal/docs", docs))

	fbApp, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: conf.Google.Project})
	if err != nil {
		log.Fatalf("Failed to create firebase app: %v", err)
	}

	fbAuth, err := fbApp.Auth(ctx)
	if err != nil {
		log.Fatalf("Failed to create firebase auth client: %v", err)
	}

	firestore, err := fbApp.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
	}
	defer firestore.Close()

	saveUser := saveuser.NewHandler(firestore)
	svc := service.New(saveUser)

	fbMW := firebaseauth.NewMiddleware(fbAuth)
	fapiPath, fapiHandler := frontendapiconnect.NewFrontendServiceHandler(svc)
	fapiHandler = fbMW(fapiHandler)
	r.Mount(fapiPath, fapiHandler)

	srv := server.NewServer(r, conf.Server)

	log.Printf("Starting server on address %v\n", srv.Addr)
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Printf("Failed to start server: %v", err)
	}
}
