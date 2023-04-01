package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetVirtualNetwork.json
func ExampleVirtualNetworksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworksClient().Get(ctx, "westus2", "myPrivateCloud", "dvportgroup-19", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetwork = armvmwarecloudsimple.VirtualNetwork{
	// 	Name: to.Ptr("Datacenter/CS-Management"),
	// 	Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualNetworks"),
	// 	Assignable: to.Ptr(false),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-19"),
	// 	Location: to.Ptr("westus2"),
	// 	Properties: &armvmwarecloudsimple.VirtualNetworkProperties{
	// 		PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
	// 	},
	// }
}
