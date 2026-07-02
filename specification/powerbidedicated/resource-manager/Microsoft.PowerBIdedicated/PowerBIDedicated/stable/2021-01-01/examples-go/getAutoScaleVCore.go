package armpowerbidedicated_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated/v2"
)

// Generated from example definition: 2021-01-01/getAutoScaleVCore.json
func ExampleAutoScaleVCoresClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerbidedicated.NewClientFactory("613192d7-503f-477a-9cfe-4efc3ee2bd60", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAutoScaleVCoresClient().Get(ctx, "TestRG", "testvcore", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpowerbidedicated.AutoScaleVCoresClientGetResponse{
	// 	AutoScaleVCore: armpowerbidedicated.AutoScaleVCore{
	// 		Name: to.Ptr("testvcore"),
	// 		Type: to.Ptr("Microsoft.PowerBIDedicated/autoScaleVCores"),
	// 		ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.PowerBIDedicated/autoScaleVCores/testvcore"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armpowerbidedicated.AutoScaleVCoreProperties{
	// 			CapacityLimit: to.Ptr[int32](10),
	// 			CapacityObjectID: to.Ptr("a28f00bd-5330-4572-88f1-fa883e074785"),
	// 			ProvisioningState: to.Ptr(armpowerbidedicated.VCoreProvisioningStateSucceeded),
	// 		},
	// 		SKU: &armpowerbidedicated.AutoScaleVCoreSKU{
	// 			Name: to.Ptr("AutoScale"),
	// 			Capacity: to.Ptr[int32](0),
	// 			Tier: to.Ptr(armpowerbidedicated.VCoreSKUTierAutoScale),
	// 		},
	// 		SystemData: &armpowerbidedicated.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T00:00:19.1234567Z"); return t}()),
	// 			CreatedBy: to.Ptr("app1"),
	// 			CreatedByType: to.Ptr(armpowerbidedicated.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T00:00:19.1234567Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("app1"),
	// 			LastModifiedByType: to.Ptr(armpowerbidedicated.CreatedByTypeApplication),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
