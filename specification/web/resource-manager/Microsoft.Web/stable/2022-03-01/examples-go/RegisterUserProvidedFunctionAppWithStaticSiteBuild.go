package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/RegisterUserProvidedFunctionAppWithStaticSiteBuild.json
func ExampleStaticSitesClient_BeginRegisterUserProvidedFunctionAppWithStaticSiteBuild() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewStaticSitesClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRegisterUserProvidedFunctionAppWithStaticSiteBuild(ctx,
		"rg",
		"testStaticSite0",
		"default",
		"testFunctionApp",
		armappservice.StaticSiteUserProvidedFunctionAppARMResource{
			Properties: &armappservice.StaticSiteUserProvidedFunctionAppARMResourceProperties{
				FunctionAppRegion:     to.Ptr("West US 2"),
				FunctionAppResourceID: to.Ptr("/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/functionRG/providers/Microsoft.Web/sites/testFunctionApp"),
			},
		},
		&armappservice.StaticSitesClientBeginRegisterUserProvidedFunctionAppWithStaticSiteBuildOptions{IsForced: to.Ptr(true)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
