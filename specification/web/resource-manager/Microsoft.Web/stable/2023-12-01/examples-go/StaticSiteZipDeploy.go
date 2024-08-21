package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21c2852d62ccc3abe9cc3800c989c6826f8363dc/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/StaticSiteZipDeploy.json
func ExampleStaticSitesClient_BeginCreateZipDeploymentForStaticSite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStaticSitesClient().BeginCreateZipDeploymentForStaticSite(ctx, "rg", "testStaticSite0", armappservice.StaticSiteZipDeploymentARMResource{
		Properties: &armappservice.StaticSiteZipDeployment{
			APIZipURL:        to.Ptr("https://teststorageaccount.net/happy-sea-15afae3e-master-81828877/api-zipdeploy.zip"),
			AppZipURL:        to.Ptr("https://teststorageaccount.net/happy-sea-15afae3e-master-81828877/app-zipdeploy.zip"),
			DeploymentTitle:  to.Ptr("Update index.html"),
			FunctionLanguage: to.Ptr("testFunctionLanguage"),
			Provider:         to.Ptr("testProvider"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
