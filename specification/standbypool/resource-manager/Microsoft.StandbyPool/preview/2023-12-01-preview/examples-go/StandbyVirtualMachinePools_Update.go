package armstandbypool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/standbypool/resource-manager/Microsoft.StandbyPool/preview/2023-12-01-preview/examples/StandbyVirtualMachinePools_Update.json
func ExampleStandbyVirtualMachinePoolsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStandbyVirtualMachinePoolsClient().Update(ctx, "rgstandbypool", "pool", armstandbypool.StandbyVirtualMachinePoolResourceUpdate{
		Properties: &armstandbypool.StandbyVirtualMachinePoolResourceUpdateProperties{
			AttachedVirtualMachineScaleSetID: to.Ptr("/subscriptions/8CC31D61-82D7-4B2B-B9DC-6B924DE7D229/resourceGroups/rgstandbypool/providers/Microsoft.Compute/virtualMachineScaleSets/myVmss"),
			ElasticityProfile: &armstandbypool.StandbyVirtualMachinePoolElasticityProfileUpdate{
				MaxReadyCapacity: to.Ptr[int64](304),
			},
			VirtualMachineState: to.Ptr(armstandbypool.VirtualMachineStateRunning),
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StandbyVirtualMachinePoolResource = armstandbypool.StandbyVirtualMachinePoolResource{
	// 	Name: to.Ptr("pool"),
	// 	Type: to.Ptr("Microsoft.StandbyPool/standbyVirtualMachinePools"),
	// 	ID: to.Ptr("/subscriptions/8CC31D61-82D7-4B2B-B9DC-6B924DE7D229/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/pool"),
	// 	SystemData: &armstandbypool.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
	// 		CreatedBy: to.Ptr("pooluser@microsoft.com"),
	// 		CreatedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("pooluser@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armstandbypool.StandbyVirtualMachinePoolResourceProperties{
	// 		AttachedVirtualMachineScaleSetID: to.Ptr("/subscriptions/8CC31D61-82D7-4B2B-B9DC-6B924DE7D229/resourceGroups/rgstandbypool/providers/Microsoft.Compute/virtualMachineScaleSets/myVmss"),
	// 		ElasticityProfile: &armstandbypool.StandbyVirtualMachinePoolElasticityProfile{
	// 			MaxReadyCapacity: to.Ptr[int64](304),
	// 		},
	// 		ProvisioningState: to.Ptr(armstandbypool.ProvisioningStateSucceeded),
	// 		VirtualMachineState: to.Ptr(armstandbypool.VirtualMachineStateRunning),
	// 	},
	// }
}
