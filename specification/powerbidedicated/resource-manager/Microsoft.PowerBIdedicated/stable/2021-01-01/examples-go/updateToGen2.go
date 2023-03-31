package armpowerbidedicated_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/updateToGen2.json
func ExampleCapacitiesClient_BeginUpdate_updateCapacityToGeneration2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerbidedicated.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCapacitiesClient().BeginUpdate(ctx, "TestRG", "azsdktest", armpowerbidedicated.DedicatedCapacityUpdateParameters{
		Properties: &armpowerbidedicated.DedicatedCapacityMutableProperties{
			Mode: to.Ptr(armpowerbidedicated.ModeGen2),
		},
		SKU: &armpowerbidedicated.CapacitySKU{
			Name: to.Ptr("A1"),
			Tier: to.Ptr(armpowerbidedicated.CapacitySKUTierPBIEAzure),
		},
		Tags: map[string]*string{
			"testKey": to.Ptr("testValue"),
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
	// res.DedicatedCapacity = armpowerbidedicated.DedicatedCapacity{
	// 	Name: to.Ptr("azsdktest"),
	// 	ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.PowerBIDedicated/capacities/azsdktest"),
	// 	Location: to.Ptr("West US"),
	// 	SystemData: &armpowerbidedicated.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T00:00:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armpowerbidedicated.IdentityTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T00:00:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armpowerbidedicated.IdentityTypeUser),
	// 	},
	// 	Tags: map[string]*string{
	// 		"testKey": to.Ptr("testValue"),
	// 	},
	// 	Properties: &armpowerbidedicated.DedicatedCapacityProperties{
	// 		Administration: &armpowerbidedicated.DedicatedCapacityAdministrators{
	// 			Members: []*string{
	// 				to.Ptr("azsdktest@microsoft.com"),
	// 				to.Ptr("azsdktest2@microsoft.com")},
	// 			},
	// 			Mode: to.Ptr(armpowerbidedicated.ModeGen2),
	// 			ProvisioningState: to.Ptr(armpowerbidedicated.CapacityProvisioningStateSucceeded),
	// 			State: to.Ptr(armpowerbidedicated.StateProvisioning),
	// 		},
	// 		SKU: &armpowerbidedicated.CapacitySKU{
	// 			Name: to.Ptr("A1"),
	// 			Tier: to.Ptr(armpowerbidedicated.CapacitySKUTierPBIEAzure),
	// 		},
	// 	}
}
