package armappplatform_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-03-01-preview/examples/BuildServiceBuilder_CreateOrUpdate.json
func ExampleBuildServiceBuilderClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armappplatform.NewBuildServiceBuilderClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<build-service-name>",
		"<builder-name>",
		armappplatform.BuilderResource{
			Properties: &armappplatform.BuilderProperties{
				BuildpackGroups: []*armappplatform.BuildpacksGroupProperties{
					{
						Name: to.Ptr("<name>"),
						Buildpacks: []*armappplatform.BuildpackProperties{
							{
								ID: to.Ptr("<id>"),
							}},
					}},
				Stack: &armappplatform.StackProperties{
					ID:      to.Ptr("<id>"),
					Version: to.Ptr("<version>"),
				},
			},
		},
		&armappplatform.BuildServiceBuilderClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
