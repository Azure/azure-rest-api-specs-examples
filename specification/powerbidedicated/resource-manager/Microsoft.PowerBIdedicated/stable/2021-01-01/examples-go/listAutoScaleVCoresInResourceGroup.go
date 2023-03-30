package armpowerbidedicated_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/listAutoScaleVCoresInResourceGroup.json
func ExampleAutoScaleVCoresClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerbidedicated.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAutoScaleVCoresClient().NewListByResourceGroupPager("TestRG", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AutoScaleVCoreListResult = armpowerbidedicated.AutoScaleVCoreListResult{
		// 	Value: []*armpowerbidedicated.AutoScaleVCore{
		// 		{
		// 			Name: to.Ptr("testvcore1"),
		// 			Type: to.Ptr("Microsoft.PowerBIDedicated/autoScaleVCores"),
		// 			ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.PowerBIDedicated/autoScaleVCores/testvcore1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"testKey": to.Ptr("testValue"),
		// 			},
		// 			Properties: &armpowerbidedicated.AutoScaleVCoreProperties{
		// 				CapacityLimit: to.Ptr[int32](10),
		// 				CapacityObjectID: to.Ptr("a28f00bd-5330-4572-88f1-fa883e074785"),
		// 				ProvisioningState: to.Ptr(armpowerbidedicated.VCoreProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armpowerbidedicated.AutoScaleVCoreSKU{
		// 				Name: to.Ptr("AutoScale"),
		// 				Capacity: to.Ptr[int32](5),
		// 				Tier: to.Ptr(armpowerbidedicated.VCoreSKUTierAutoScale),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testvcore2"),
		// 			Type: to.Ptr("Microsoft.PowerBIDedicated/autoScaleVCores"),
		// 			ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.PowerBIDedicated/autoScaleVCores/testvcore2"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"testKey": to.Ptr("testValue"),
		// 			},
		// 			Properties: &armpowerbidedicated.AutoScaleVCoreProperties{
		// 				CapacityLimit: to.Ptr[int32](10),
		// 				CapacityObjectID: to.Ptr("f59f226b-a313-4266-8768-a45961c2edbb"),
		// 				ProvisioningState: to.Ptr(armpowerbidedicated.VCoreProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armpowerbidedicated.AutoScaleVCoreSKU{
		// 				Name: to.Ptr("AutoScale"),
		// 				Capacity: to.Ptr[int32](0),
		// 				Tier: to.Ptr(armpowerbidedicated.VCoreSKUTierAutoScale),
		// 			},
		// 	}},
		// }
	}
}
