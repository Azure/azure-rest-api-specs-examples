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
					DiagnosticLevel:           armazurestackhci.DiagnosticLevel("Basic").ToPtr(),
					WindowsServerSubscription: armazurestackhci.WindowsServerSubscription("Enabled").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.ClustersClientUpdateResult)
}
