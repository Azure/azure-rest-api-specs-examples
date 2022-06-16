package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci"
)

// x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2021-09-01/examples/CreateCluster.json
func ExampleClustersClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurestackhci.NewClustersClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armazurestackhci.Cluster{
			Location: to.StringPtr("<location>"),
			Properties: &armazurestackhci.ClusterProperties{
				AADClientID:             to.StringPtr("<aadclient-id>"),
				AADTenantID:             to.StringPtr("<aadtenant-id>"),
				CloudManagementEndpoint: to.StringPtr("<cloud-management-endpoint>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClustersClientCreateResult)
}
