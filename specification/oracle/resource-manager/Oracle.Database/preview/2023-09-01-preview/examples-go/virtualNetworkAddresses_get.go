package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/520e274d7d95fc6d1002dd3c1fcaf8d55d27f63e/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/virtualNetworkAddresses_get.json
func ExampleVirtualNetworkAddressesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkAddressesClient().Get(ctx, "rg000", "cluster1", "hostname1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkAddress = armoracledatabase.VirtualNetworkAddress{
	// 	Type: to.Ptr("Oracle.Database/cloudVmClusters/virtualNetworkAddresses"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudVmClusters/cluster1/virtualNetworkAddresses/hostname1"),
	// 	Properties: &armoracledatabase.VirtualNetworkAddressProperties{
	// 		Domain: to.Ptr("domain1"),
	// 		IPAddress: to.Ptr("192.168.0.1"),
	// 		LifecycleDetails: to.Ptr("success"),
	// 		LifecycleState: to.Ptr(armoracledatabase.VirtualNetworkAddressLifecycleStateAvailable),
	// 		Ocid: to.Ptr("ocid1..aaaaaa"),
	// 		ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 		TimeAssigned: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-22T00:27:02.119Z"); return t}()),
	// 		VMOcid: to.Ptr("ocid1..aaaa"),
	// 	},
	// }
}
