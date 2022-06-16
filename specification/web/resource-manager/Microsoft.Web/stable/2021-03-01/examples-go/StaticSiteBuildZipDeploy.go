package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/StaticSiteBuildZipDeploy.json
func ExampleStaticSitesClient_BeginCreateZipDeploymentForStaticSiteBuild() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewStaticSitesClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateZipDeploymentForStaticSiteBuild(ctx,
		"rg",
		"testStaticSite0",
		"12",
		armappservice.StaticSiteZipDeploymentARMResource{
			Properties: &armappservice.StaticSiteZipDeployment{
				APIZipURL:        to.Ptr("https://teststorageaccount.net/happy-sea-15afae3e-master-81828877/api-zipdeploy.zip"),
				AppZipURL:        to.Ptr("https://teststorageaccount.net/happy-sea-15afae3e-master-81828877/app-zipdeploy.zip"),
				DeploymentTitle:  to.Ptr("Update index.html"),
				FunctionLanguage: to.Ptr("testFunctionLanguage"),
				Provider:         to.Ptr("testProvider"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
