package armstoragepool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/Skus_List.json
func ExampleResourceSKUsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragepool.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceSKUsClient().NewListPager("eastus", nil)
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
		// page.ResourceSKUListResult = armstoragepool.ResourceSKUListResult{
		// 	Value: []*armstoragepool.ResourceSKUInfo{
		// 		{
		// 			Name: to.Ptr("Standard_V1"),
		// 			APIVersion: to.Ptr("2021-08-01"),
		// 			Capabilities: []*armstoragepool.ResourceSKUCapability{
		// 				{
		// 					Name: to.Ptr("MaxNumberOfDisks"),
		// 					Value: to.Ptr("1"),
		// 			}},
		// 			LocationInfo: &armstoragepool.ResourceSKULocationInfo{
		// 				Location: to.Ptr("eastus"),
		// 				ZoneDetails: []*armstoragepool.ResourceSKUZoneDetails{
		// 					{
		// 						Name: []*string{
		// 							to.Ptr("2")},
		// 							Capabilities: []*armstoragepool.ResourceSKUCapability{
		// 								{
		// 									Name: to.Ptr("DiskPool.CapabilityA"),
		// 									Value: to.Ptr("True"),
		// 							}},
		// 						},
		// 						{
		// 							Name: []*string{
		// 								to.Ptr("1")},
		// 								Capabilities: []*armstoragepool.ResourceSKUCapability{
		// 									{
		// 										Name: to.Ptr("DiskPool.CapabilityA"),
		// 										Value: to.Ptr("True"),
		// 									},
		// 									{
		// 										Name: to.Ptr("DiskPool.CapabilityB"),
		// 										Value: to.Ptr("True"),
		// 								}},
		// 						}},
		// 						Zones: []*string{
		// 							to.Ptr("2"),
		// 							to.Ptr("1")},
		// 						},
		// 						ResourceType: to.Ptr("diskPools"),
		// 						Restrictions: []*armstoragepool.ResourceSKURestrictions{
		// 							{
		// 								Type: to.Ptr(armstoragepool.ResourceSKURestrictionsTypeLocation),
		// 								ReasonCode: to.Ptr(armstoragepool.ResourceSKURestrictionsReasonCodeNotAvailableForSubscription),
		// 								RestrictionInfo: &armstoragepool.ResourceSKURestrictionInfo{
		// 									Locations: []*string{
		// 										to.Ptr("FranceSouth")},
		// 									},
		// 									Values: []*string{
		// 										to.Ptr("FranceSouth")},
		// 									},
		// 									{
		// 										Type: to.Ptr(armstoragepool.ResourceSKURestrictionsTypeZone),
		// 										ReasonCode: to.Ptr(armstoragepool.ResourceSKURestrictionsReasonCodeNotAvailableForSubscription),
		// 										RestrictionInfo: &armstoragepool.ResourceSKURestrictionInfo{
		// 											Locations: []*string{
		// 												to.Ptr("FranceCentral")},
		// 												Zones: []*string{
		// 													to.Ptr("2"),
		// 													to.Ptr("3"),
		// 													to.Ptr("1")},
		// 												},
		// 												Values: []*string{
		// 													to.Ptr("FranceCentral")},
		// 											}},
		// 											Tier: to.Ptr("Standard"),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_V1"),
		// 											Capabilities: []*armstoragepool.ResourceSKUCapability{
		// 												{
		// 													Name: to.Ptr("MaxNumberOfDisks"),
		// 													Value: to.Ptr("1"),
		// 											}},
		// 											LocationInfo: &armstoragepool.ResourceSKULocationInfo{
		// 												Location: to.Ptr("eastus"),
		// 												ZoneDetails: []*armstoragepool.ResourceSKUZoneDetails{
		// 													{
		// 														Name: []*string{
		// 															to.Ptr("2")},
		// 															Capabilities: []*armstoragepool.ResourceSKUCapability{
		// 																{
		// 																	Name: to.Ptr("UltraSSDAvailable"),
		// 																	Value: to.Ptr("True"),
		// 															}},
		// 														},
		// 														{
		// 															Name: []*string{
		// 																to.Ptr("1")},
		// 																Capabilities: []*armstoragepool.ResourceSKUCapability{
		// 																	{
		// 																		Name: to.Ptr("UltraSSDAvailable"),
		// 																		Value: to.Ptr("True"),
		// 																}},
		// 														}},
		// 														Zones: []*string{
		// 															to.Ptr("2"),
		// 															to.Ptr("1")},
		// 														},
		// 														ResourceType: to.Ptr("diskPools"),
		// 														Tier: to.Ptr("Standard"),
		// 												}},
		// 											}
	}
}
