package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NspLinkPut.json
func ExampleSecurityPerimeterLinksClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecurityPerimeterLinksClient().CreateOrUpdate(ctx, "rg1", "nsp1", "link1", armnetwork.NspLink{
		Properties: &armnetwork.NspLinkProperties{
			AutoApprovedRemotePerimeterResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp2"),
			LocalInboundProfiles: []*string{
				to.Ptr("*"),
			},
			RemoteInboundProfiles: []*string{
				to.Ptr("*"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.SecurityPerimeterLinksClientCreateOrUpdateResponse{
	// 	NspLink: armnetwork.NspLink{
	// 		Name: to.Ptr("link1"),
	// 		Type: to.Ptr("Microsoft.Network/networkSecurityPerimeters/links"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/links/link1"),
	// 		Properties: &armnetwork.NspLinkProperties{
	// 			Description: to.Ptr("Auto Approved"),
	// 			AutoApprovedRemotePerimeterResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp2"),
	// 			LocalInboundProfiles: []*string{
	// 				to.Ptr("*"),
	// 			},
	// 			LocalOutboundProfiles: []*string{
	// 				to.Ptr("*"),
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.NspLinkProvisioningStateSucceeded),
	// 			RemoteInboundProfiles: []*string{
	// 				to.Ptr("*"),
	// 			},
	// 			RemoteOutboundProfiles: []*string{
	// 				to.Ptr("*"),
	// 			},
	// 			RemotePerimeterGUID: to.Ptr("guid"),
	// 			RemotePerimeterLocation: to.Ptr("westus2"),
	// 			Status: to.Ptr(armnetwork.NspLinkStatusApproved),
	// 		},
	// 		SystemData: &armnetwork.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.3446713Z"); return t}()),
	// 			CreatedBy: to.Ptr("user"),
	// 			CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.3446713Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user"),
	// 			LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
