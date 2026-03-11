package armdisconnectedoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/disconnectedoperations/armdisconnectedoperations"
)

// Generated from example definition: 2026-03-15/HardwareSettings_CreateOrUpdate_MaximumSet_Gen.json
func ExampleHardwareSettingsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdisconnectedoperations.NewClientFactory("AFEEE483-435F-4E9C-8742-4B550746CD70", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewHardwareSettingsClient().BeginCreateOrUpdate(ctx, "rgdisconnectedOperations", "demo-resource", "default", armdisconnectedoperations.HardwareSetting{
		Properties: &armdisconnectedoperations.HardwareSettingProperties{
			TotalCores:               to.Ptr[int32](200),
			DiskSpaceInGb:            to.Ptr[int32](1024),
			MemoryInGb:               to.Ptr[int32](64),
			Oem:                      to.Ptr("Contoso"),
			HardwareSKU:              to.Ptr("MC-760"),
			Nodes:                    to.Ptr[int32](3),
			VersionAtRegistration:    to.Ptr("2411.2"),
			SolutionBuilderExtension: to.Ptr("xyz"),
			DeviceID:                 to.Ptr("663ee8a3-4ea8-48ec-8810-b1f8b86cb270"),
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
	// res = armdisconnectedoperations.HardwareSettingsClientCreateOrUpdateResponse{
	// 	HardwareSetting: &armdisconnectedoperations.HardwareSetting{
	// 		Properties: &armdisconnectedoperations.HardwareSettingProperties{
	// 			ProvisioningState: to.Ptr(armdisconnectedoperations.ResourceProvisioningStateSucceeded),
	// 			TotalCores: to.Ptr[int32](200),
	// 			DiskSpaceInGb: to.Ptr[int32](1024),
	// 			MemoryInGb: to.Ptr[int32](64),
	// 			Oem: to.Ptr("Contoso"),
	// 			HardwareSKU: to.Ptr("MC-760"),
	// 			Nodes: to.Ptr[int32](3),
	// 			VersionAtRegistration: to.Ptr("2411.2"),
	// 			SolutionBuilderExtension: to.Ptr("xyz"),
	// 			DeviceID: to.Ptr("663ee8a3-4ea8-48ec-8810-b1f8b86cb270"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/AFEEE483-435F-4E9C-8742-4B550746CD70/resourceGroups/rgdisconnectedOperations/providers/Microsoft.Edge/disconnectedOperations/demo-resource"),
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.Edge/disconnectedOperations/hardwareSettings"),
	// 		SystemData: &armdisconnectedoperations.SystemData{
	// 			CreatedBy: to.Ptr("kisu"),
	// 			CreatedByType: to.Ptr(armdisconnectedoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-15T22:59:45.585Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("alaunn"),
	// 			LastModifiedByType: to.Ptr(armdisconnectedoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-15T22:59:45.585Z"); return t}()),
	// 		},
	// 	},
	// }
}
