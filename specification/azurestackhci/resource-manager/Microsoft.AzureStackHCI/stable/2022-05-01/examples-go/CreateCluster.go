package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/CreateCluster.json
func ExampleClustersClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armazurestackhci.NewClustersClient("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"test-rg",
		"myCluster",
		armazurestackhci.Cluster{
			Location: to.Ptr("East US"),
			Properties: &armazurestackhci.ClusterProperties{
				AADClientID:             to.Ptr("24a6e53d-04e5-44d2-b7cc-1b732a847dfc"),
				AADTenantID:             to.Ptr("7e589cc1-a8b6-4dff-91bd-5ec0fa18db94"),
				CloudManagementEndpoint: to.Ptr("https://98294836-31be-4668-aeae-698667faf99b.waconazure.com"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
