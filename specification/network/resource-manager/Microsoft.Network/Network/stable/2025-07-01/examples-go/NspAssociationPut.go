package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NspAssociationPut.json
func ExampleSecurityPerimeterAssociationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSecurityPerimeterAssociationsClient().BeginCreateOrUpdate(ctx, "rg1", "nsp1", "association1", armnetwork.NspAssociation{
		Properties: &armnetwork.NspAssociationProperties{
			AccessMode: to.Ptr(armnetwork.AssociationAccessModeEnforced),
			PrivateLinkResource: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/{paasSubscriptionId}/resourceGroups/{paasResourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}"),
			},
			Profile: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/profiles/{profileName}"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.SecurityPerimeterAssociationsClientCreateOrUpdateResponse{
	// 	NspAssociation: armnetwork.NspAssociation{
	// 		Name: to.Ptr("association1"),
	// 		Type: to.Ptr("Microsoft.Network/networkSecurityPerimeters/resourceAssociations"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/resourceAssociations/association1"),
	// 		Properties: &armnetwork.NspAssociationProperties{
	// 			AccessMode: to.Ptr(armnetwork.AssociationAccessModeEnforced),
	// 			HasProvisioningIssues: to.Ptr("no"),
	// 			PrivateLinkResource: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/{paasSubscriptionId}/resourceGroups/{paasResourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}"),
	// 			},
	// 			Profile: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/profiles/{profileName}"),
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.NspProvisioningStateSucceeded),
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
