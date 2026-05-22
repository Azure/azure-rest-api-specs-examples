package armpowerbidedicated_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated/v2"
)

// Generated from example definition: 2021-01-01/listCapacitiesInSubscription.json
func ExampleCapacitiesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerbidedicated.NewClientFactory("613192d7-503f-477a-9cfe-4efc3ee2bd60", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCapacitiesClient().NewListPager(nil)
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
		// page = armpowerbidedicated.CapacitiesClientListResponse{
		// 	DedicatedCapacities: armpowerbidedicated.DedicatedCapacities{
		// 		Value: []*armpowerbidedicated.DedicatedCapacity{
		// 			{
		// 				Name: to.Ptr("azsdktest"),
		// 				ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/providers/Microsoft.PowerBIDedicated/capacities/azsdktest"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armpowerbidedicated.DedicatedCapacityProperties{
		// 					Administration: &armpowerbidedicated.DedicatedCapacityAdministrators{
		// 						Members: []*string{
		// 							to.Ptr("azsdktest@microsoft.com"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armpowerbidedicated.CapacityProvisioningStateProvisioning),
		// 					State: to.Ptr(armpowerbidedicated.StateProvisioning),
		// 				},
		// 				SKU: &armpowerbidedicated.CapacitySKU{
		// 					Name: to.Ptr("A1"),
		// 					Tier: to.Ptr(armpowerbidedicated.CapacitySKUTierPBIEAzure),
		// 				},
		// 				Tags: map[string]*string{
		// 					"testKey": to.Ptr("testValue"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("azsdktest2"),
		// 				ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/providers/Microsoft.PowerBIDedicated/capacities/azsdktest"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armpowerbidedicated.DedicatedCapacityProperties{
		// 					Administration: &armpowerbidedicated.DedicatedCapacityAdministrators{
		// 						Members: []*string{
		// 							to.Ptr("azsdktest@microsoft.com"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armpowerbidedicated.CapacityProvisioningStateProvisioning),
		// 					State: to.Ptr(armpowerbidedicated.StateProvisioning),
		// 				},
		// 				SKU: &armpowerbidedicated.CapacitySKU{
		// 					Name: to.Ptr("A2"),
		// 					Tier: to.Ptr(armpowerbidedicated.CapacitySKUTierPBIEAzure),
		// 				},
		// 				Tags: map[string]*string{
		// 					"testKey": to.Ptr("testValue"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
