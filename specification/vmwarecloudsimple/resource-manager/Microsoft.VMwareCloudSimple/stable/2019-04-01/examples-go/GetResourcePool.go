package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetResourcePool.json
func ExampleResourcePoolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourcePoolsClient().Get(ctx, "westus2", "myPrivateCloud", "resgroup-26", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ResourcePool = armvmwarecloudsimple.ResourcePool{
	// 	Name: to.Ptr("Workload"),
	// 	Type: to.Ptr("Microsoft.VMwareCloudSimple/resourcePools"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/resourcePools/resgroup-26"),
	// 	Location: to.Ptr("westus2"),
	// 	PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
	// 	Properties: &armvmwarecloudsimple.ResourcePoolProperties{
	// 		FullName: to.Ptr("myPrivateCloud/Datacenter/Cluster/Workload"),
	// 	},
	// }
}
