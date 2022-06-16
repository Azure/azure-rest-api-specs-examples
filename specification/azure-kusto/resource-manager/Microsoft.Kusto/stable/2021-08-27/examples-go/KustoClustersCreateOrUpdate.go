package armkusto_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClustersCreateOrUpdate.json
func ExampleClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armkusto.Cluster{
			Location: to.StringPtr("<location>"),
			Identity: &armkusto.Identity{
				Type: armkusto.IdentityType("SystemAssigned").ToPtr(),
			},
			Properties: &armkusto.ClusterProperties{
				AllowedIPRangeList: []*string{
					to.StringPtr("0.0.0.0/0")},
				EnableAutoStop:         to.BoolPtr(true),
				EnableDoubleEncryption: to.BoolPtr(false),
				EnablePurge:            to.BoolPtr(true),
				EnableStreamingIngest:  to.BoolPtr(true),
				PublicNetworkAccess:    armkusto.PublicNetworkAccess("Enabled").ToPtr(),
			},
			SKU: &armkusto.AzureSKU{
				Name:     armkusto.AzureSKUName("Standard_L8s").ToPtr(),
				Capacity: to.Int32Ptr(2),
				Tier:     armkusto.AzureSKUTier("Standard").ToPtr(),
			},
		},
		&armkusto.ClustersClientBeginCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClustersClientCreateOrUpdateResult)
}
