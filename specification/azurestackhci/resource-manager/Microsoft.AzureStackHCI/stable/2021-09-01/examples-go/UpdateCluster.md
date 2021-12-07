Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazurestackhci%2Farmazurestackhci%2Fv0.1.0/sdk/resourcemanager/azurestackhci/armazurestackhci/README.md) on how to add the SDK to your project and authenticate.

```go
package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci"
)

// x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2021-09-01/examples/UpdateCluster.json
func ExampleClustersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurestackhci.NewClustersClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armazurestackhci.ClusterPatch{
			Properties: &armazurestackhci.ClusterPatchProperties{
				CloudManagementEndpoint: to.StringPtr("<cloud-management-endpoint>"),
				DesiredProperties: &armazurestackhci.ClusterDesiredProperties{
					DiagnosticLevel:           armazurestackhci.DiagnosticLevelBasic.ToPtr(),
					WindowsServerSubscription: armazurestackhci.WindowsServerSubscriptionEnabled.ToPtr(),
				},
			},
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Cluster.ID: %s\n", *res.ID)
}
```
