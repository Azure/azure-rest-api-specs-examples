package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoClustersCreateOrUpdate.json
func ExampleClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkusto.NewClustersClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"kustorptest",
		"kustoCluster",
		armkusto.Cluster{
			Location: to.Ptr("westus"),
			Identity: &armkusto.Identity{
				Type: to.Ptr(armkusto.IdentityTypeSystemAssigned),
			},
			Properties: &armkusto.ClusterProperties{
				AllowedIPRangeList: []*string{
					to.Ptr("0.0.0.0/0")},
				EnableAutoStop:         to.Ptr(true),
				EnableDoubleEncryption: to.Ptr(false),
				EnablePurge:            to.Ptr(true),
				EnableStreamingIngest:  to.Ptr(true),
				PublicIPType:           to.Ptr(armkusto.PublicIPTypeDualStack),
				PublicNetworkAccess:    to.Ptr(armkusto.PublicNetworkAccessEnabled),
			},
			SKU: &armkusto.AzureSKU{
				Name:     to.Ptr(armkusto.AzureSKUNameStandardL8S),
				Capacity: to.Ptr[int32](2),
				Tier:     to.Ptr(armkusto.AzureSKUTierStandard),
			},
		},
		&armkusto.ClustersClientBeginCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
