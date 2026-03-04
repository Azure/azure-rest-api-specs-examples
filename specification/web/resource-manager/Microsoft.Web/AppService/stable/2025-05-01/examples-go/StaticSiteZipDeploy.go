package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v6"
)

// Generated from example definition: 2025-05-01/StaticSiteZipDeploy.json
func ExampleStaticSitesClient_BeginCreateZipDeploymentForStaticSite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStaticSitesClient().BeginCreateZipDeploymentForStaticSite(ctx, "rg", "testStaticSite0", armappservice.StaticSiteZipDeploymentARMResource{
		Properties: &armappservice.StaticSiteZipDeployment{
			APIZipURL:        to.Ptr("https://[examplestorageaccount].com/happy-sea-15afae3e-master-81828877/api-zipdeploy.zip"),
			AppZipURL:        to.Ptr("https://[examplestorageaccount].com/happy-sea-15afae3e-master-81828877/app-zipdeploy.zip"),
			DeploymentTitle:  to.Ptr("Update index.html"),
			FunctionLanguage: to.Ptr("testFunctionLanguage"),
			Provider:         to.Ptr("testProvider"),
		},
	}, nil)
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
	// res = armappservice.StaticSitesClientCreateZipDeploymentForStaticSiteResponse{
	// }
}
