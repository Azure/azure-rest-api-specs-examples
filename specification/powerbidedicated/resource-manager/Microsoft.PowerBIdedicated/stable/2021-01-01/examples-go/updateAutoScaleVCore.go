package armpowerbidedicated_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/updateAutoScaleVCore.json
func ExampleAutoScaleVCoresClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpowerbidedicated.NewAutoScaleVCoresClient("613192d7-503f-477a-9cfe-4efc3ee2bd60", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"TestRG",
		"testvcore",
		armpowerbidedicated.AutoScaleVCoreUpdateParameters{
			Properties: &armpowerbidedicated.AutoScaleVCoreMutableProperties{
				CapacityLimit: to.Ptr[int32](20),
			},
			SKU: &armpowerbidedicated.AutoScaleVCoreSKU{
				Name:     to.Ptr("AutoScale"),
				Capacity: to.Ptr[int32](0),
				Tier:     to.Ptr(armpowerbidedicated.VCoreSKUTierAutoScale),
			},
			Tags: map[string]*string{
				"testKey": to.Ptr("testValue"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
