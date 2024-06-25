package armscvmm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/Clouds_Get_MaximumSet_Gen.json
func ExampleCloudsClient_Get_cloudsGetMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCloudsClient().Get(ctx, "rgscvmm", "_", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Cloud = armscvmm.Cloud{
	// 	Name: to.Ptr("wwcwalpiufsfbnydxpr"),
	// 	Type: to.Ptr("qnaaimszbuokldohwrdfuiitpy"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/clouds/cloudResourceName"),
	// 	SystemData: &armscvmm.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.094Z"); return t}()),
	// 		CreatedBy: to.Ptr("p"),
	// 		CreatedByType: to.Ptr(armscvmm.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.095Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("goxcwpyyqlxndquly"),
	// 		LastModifiedByType: to.Ptr(armscvmm.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("khwsdmaxfhmbu"),
	// 	Tags: map[string]*string{
	// 		"key4295": to.Ptr("wngosgcbdifaxdobufuuqxtho"),
	// 	},
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/customLocationName"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armscvmm.CloudProperties{
	// 		CloudCapacity: &armscvmm.CloudCapacity{
	// 			CPUCount: to.Ptr[int64](4),
	// 			MemoryMB: to.Ptr[int64](19),
	// 			VMCount: to.Ptr[int64](28),
	// 		},
	// 		CloudName: to.Ptr("menarjsplhcqvnkjdwieroir"),
	// 		InventoryItemID: to.Ptr("qjd"),
	// 		ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 		StorageQosPolicies: []*armscvmm.StorageQosPolicy{
	// 			{
	// 				Name: to.Ptr("hvqcentnbwcunxhzfavyewhwlo"),
	// 				BandwidthLimit: to.Ptr[int64](26),
	// 				ID: to.Ptr("oclhgkydaw"),
	// 				IopsMaximum: to.Ptr[int64](6),
	// 				IopsMinimum: to.Ptr[int64](25),
	// 				PolicyID: to.Ptr("lvcylbmxrqjgarvhfny"),
	// 		}},
	// 		UUID: to.Ptr("12345678-1234-1234-1234-12345678abcd"),
	// 		VmmServerID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/vmmServers/vmmServerName"),
	// 	},
	// }
}
