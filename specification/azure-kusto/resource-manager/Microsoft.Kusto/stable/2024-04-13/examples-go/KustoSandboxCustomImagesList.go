package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8adce17dc500f338f86f18af30aac61dcb71c5f/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoSandboxCustomImagesList.json
func ExampleSandboxCustomImagesClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSandboxCustomImagesClient().NewListByClusterPager("kustorptest", "kustoCluster", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SandboxCustomImagesListResult = armkusto.SandboxCustomImagesListResult{
		// 	Value: []*armkusto.SandboxCustomImage{
		// 		{
		// 			Name: to.Ptr("kustoCluster/customImage7"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/SandboxCustomImages"),
		// 			ID: to.Ptr("/subscriptions/f12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/SandboxCustomImages/customImage7"),
		// 			Properties: &armkusto.SandboxCustomImageProperties{
		// 				LanguageVersion: to.Ptr("3.10.8"),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 				RequirementsFileContent: to.Ptr("Requests"),
		// 				Language: to.Ptr(armkusto.LanguagePython),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("kustoCluster/customImage8"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/SandboxCustomImages"),
		// 			ID: to.Ptr("/subscriptions/f12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/SandboxCustomImages/customImage8"),
		// 			Properties: &armkusto.SandboxCustomImageProperties{
		// 				LanguageVersion: to.Ptr("3.9.7"),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 				RequirementsFileContent: to.Ptr("PyTorch"),
		// 				Language: to.Ptr(armkusto.LanguagePython),
		// 			},
		// 	}},
		// }
	}
}
