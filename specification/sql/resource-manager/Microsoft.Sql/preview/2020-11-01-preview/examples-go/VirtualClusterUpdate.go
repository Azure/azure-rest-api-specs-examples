package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/VirtualClusterUpdate.json
func ExampleVirtualClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualClustersClient().BeginUpdate(ctx, "testrg", "vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2", armsql.VirtualClusterUpdate{
		Properties: &armsql.VirtualClusterProperties{
			MaintenanceConfigurationID: to.Ptr("/subscriptions/ab0e51c0-83c0-4380-8ae9-025516df392f/resourceGroups/Federation/providers/Microsoft.Maintenance/maintenanceConfigurations/MiPolicy1"),
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
	// res.VirtualCluster = armsql.VirtualCluster{
	// 	Name: to.Ptr("vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2"),
	// 	Type: to.Ptr("Microsoft.Sql/virtualClusters"),
	// 	ID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Sql/virtualClusters/vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2"),
	// 	Location: to.Ptr("onebox"),
	// 	Properties: &armsql.VirtualClusterProperties{
	// 		ChildResources: []*string{
	// 			to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance1"),
	// 			to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance2")},
	// 			MaintenanceConfigurationID: to.Ptr("/subscriptions/ab0e51c0-83c0-4380-8ae9-025516df392f/resourceGroups/Federation/providers/Microsoft.Maintenance/maintenanceConfigurations/MiPolicy1"),
	// 			SubnetID: to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 		},
	// 	}
}
