```go
package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/UpdateCluster.json
func ExampleClustersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armazurestackhci.NewClustersClient("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"test-rg",
		"myCluster",
		armazurestackhci.ClusterPatch{
			Properties: &armazurestackhci.ClusterPatchProperties{
				CloudManagementEndpoint: to.Ptr("https://98294836-31be-4668-aeae-698667faf99b.waconazure.com"),
				DesiredProperties: &armazurestackhci.ClusterDesiredProperties{
					DiagnosticLevel:           to.Ptr(armazurestackhci.DiagnosticLevelBasic),
					WindowsServerSubscription: to.Ptr(armazurestackhci.WindowsServerSubscriptionEnabled),
				},
			},
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazurestackhci%2Farmazurestackhci%2Fv1.0.0/sdk/resourcemanager/azurestackhci/armazurestackhci/README.md) on how to add the SDK to your project and authenticate.
