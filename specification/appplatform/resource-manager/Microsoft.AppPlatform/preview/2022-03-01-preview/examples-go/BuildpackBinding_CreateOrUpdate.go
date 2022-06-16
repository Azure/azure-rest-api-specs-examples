package armappplatform_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-03-01-preview/examples/BuildpackBinding_CreateOrUpdate.json
func ExampleBuildpackBindingClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armappplatform.NewBuildpackBindingClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<build-service-name>",
		"<builder-name>",
		"<buildpack-binding-name>",
		armappplatform.BuildpackBindingResource{
			Properties: &armappplatform.BuildpackBindingProperties{
				BindingType: to.Ptr(armappplatform.BindingTypeApplicationInsights),
				LaunchProperties: &armappplatform.BuildpackBindingLaunchProperties{
					Properties: map[string]*string{
						"abc":           to.Ptr("def"),
						"any-string":    to.Ptr("any-string"),
						"sampling-rate": to.Ptr("12.0"),
					},
					Secrets: map[string]*string{
						"connection-string": to.Ptr("XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXX;XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXXXXXXXX"),
					},
				},
			},
		},
		&armappplatform.BuildpackBindingClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
