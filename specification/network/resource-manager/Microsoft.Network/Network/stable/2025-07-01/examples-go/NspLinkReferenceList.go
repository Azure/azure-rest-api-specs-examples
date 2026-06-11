package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NspLinkReferenceList.json
func ExampleSecurityPerimeterLinkReferencesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecurityPerimeterLinkReferencesClient().NewListPager("rg1", "nsp2", nil)
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
		// page = armnetwork.SecurityPerimeterLinkReferencesClientListResponse{
		// 	NspLinkReferenceListResult: armnetwork.NspLinkReferenceListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp2/linkreferences?api-version=2025-07-01&$skipToken=10"),
		// 		Value: []*armnetwork.NspLinkReference{
		// 			{
		// 				Name: to.Ptr("link1-guid"),
		// 				Type: to.Ptr("Microsoft.Network/networkSecurityPerimeters/linkreferences"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp2/linkreferences/link1-guid"),
		// 				Properties: &armnetwork.NspLinkReferenceProperties{
		// 					Description: to.Ptr("Auto Approved"),
		// 					LocalInboundProfiles: []*string{
		// 						to.Ptr("*"),
		// 					},
		// 					LocalOutboundProfiles: []*string{
		// 						to.Ptr("*"),
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.NspLinkProvisioningStateSucceeded),
		// 					RemoteInboundProfiles: []*string{
		// 						to.Ptr("*"),
		// 					},
		// 					RemoteOutboundProfiles: []*string{
		// 						to.Ptr("*"),
		// 					},
		// 					RemotePerimeterGUID: to.Ptr("guid"),
		// 					RemotePerimeterLocation: to.Ptr("westus2"),
		// 					RemotePerimeterResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1"),
		// 					Status: to.Ptr(armnetwork.NspLinkStatusApproved),
		// 				},
		// 				SystemData: &armnetwork.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.3446713Z"); return t}()),
		// 					CreatedBy: to.Ptr("user"),
		// 					CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.3446713Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user"),
		// 					LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
