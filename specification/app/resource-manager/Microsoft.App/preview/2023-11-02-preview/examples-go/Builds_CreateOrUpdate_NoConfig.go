package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/Builds_CreateOrUpdate_NoConfig.json
func ExampleBuildsClient_BeginCreateOrUpdate_buildsCreateOrUpdateNoConfig() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBuildsClient().BeginCreateOrUpdate(ctx, "rg", "testBuilder", "testBuild", armappcontainers.BuildResource{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BuildResource = armappcontainers.BuildResource{
	// 	Name: to.Ptr("testBuild"),
	// 	Type: to.Ptr("Microsoft.App/builders/builds"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/Microsoft.App/builders/testBuilder/builds/testBuild"),
	// 	SystemData: &armappcontainers.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-11T11:05:51.494Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample@microsoft.com"),
	// 		CreatedByType: to.Ptr(armappcontainers.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-11T11:05:51.494Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armappcontainers.CreatedByTypeUser),
	// 	},
	// 	Properties: &armappcontainers.BuildProperties{
	// 		BuildStatus: to.Ptr(armappcontainers.BuildStatusInProgress),
	// 		LogStreamEndpoint: to.Ptr("https://foo.azurecontainerapps.dev/logstream"),
	// 		ProvisioningState: to.Ptr(armappcontainers.BuildProvisioningStateSucceeded),
	// 		TokenEndpoint: to.Ptr("https://management.azure.com/subscriptions/{subscription-id}/resourcegroups/{rg-id}/Microsoft.App/builders/testBuilder/builds/testBuild/listAuthToken"),
	// 		UploadEndpoint: to.Ptr("https://foo.azurecontainerapps.dev/upload"),
	// 	},
	// }
}
