package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoSandboxCustomImageUpdate.json
func ExampleSandboxCustomImagesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSandboxCustomImagesClient().BeginUpdate(ctx, "kustorptest", "kustoCluster", "customImage8", armkusto.SandboxCustomImage{
		Properties: &armkusto.SandboxCustomImageProperties{
			LanguageVersion:         to.Ptr("3.10.8"),
			RequirementsFileContent: to.Ptr("Requests"),
			Language:                to.Ptr(armkusto.LanguagePython),
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
	// res.SandboxCustomImage = armkusto.SandboxCustomImage{
	// 	Name: to.Ptr("kustoCluster/customImage8"),
	// 	Type: to.Ptr("Microsoft.Kusto/Clusters/SandboxCustomImages"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/SandboxCustomImages/customImage8"),
	// 	Properties: &armkusto.SandboxCustomImageProperties{
	// 		LanguageVersion: to.Ptr("3.10.8"),
	// 		ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
	// 		RequirementsFileContent: to.Ptr("Requests"),
	// 		Language: to.Ptr(armkusto.LanguagePython),
	// 	},
	// }
}
