package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-03-01-preview/examples/BuildService_CreateOrUpdateBuild.json
func ExampleBuildServiceClient_CreateOrUpdateBuild() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armappplatform.NewBuildServiceClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdateBuild(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<build-service-name>",
		"<build-name>",
		armappplatform.Build{
			Properties: &armappplatform.BuildProperties{
				AgentPool: to.Ptr("<agent-pool>"),
				Builder:   to.Ptr("<builder>"),
				Env: map[string]*string{
					"environmentVariable": to.Ptr("test"),
				},
				RelativePath: to.Ptr("<relative-path>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
