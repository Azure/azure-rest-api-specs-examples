package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NspAssociationList.json
func ExampleSecurityPerimeterAssociationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecurityPerimeterAssociationsClient().NewListPager("rg1", "nsp1", nil)
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
		// page = armnetwork.SecurityPerimeterAssociationsClientListResponse{
		// 	NspAssociationsListResult: armnetwork.NspAssociationsListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/resourceAssociations?api-version=2025-07-01&$skipToken=10"),
		// 		Value: []*armnetwork.NspAssociation{
		// 			{
		// 				Name: to.Ptr("association1"),
		// 				Type: to.Ptr("Microsoft.Network/networkSecurityPerimeters/resourceAssociations"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/resourceAssociations/association1"),
		// 				Properties: &armnetwork.NspAssociationProperties{
		// 					AccessMode: to.Ptr(armnetwork.AssociationAccessModeEnforced),
		// 					HasProvisioningIssues: to.Ptr("no"),
		// 					PrivateLinkResource: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/{paasSubscriptionId}/resourceGroups/{paasResourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}"),
		// 					},
		// 					Profile: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/profiles/{profileName}"),
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.NspProvisioningStateSucceeded),
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
		// 			{
		// 				Name: to.Ptr("association2"),
		// 				Type: to.Ptr("Microsoft.Network/networkSecurityPerimeters/resourceAssociations"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/resourceAssociations/association2"),
		// 				Properties: &armnetwork.NspAssociationProperties{
		// 					AccessMode: to.Ptr(armnetwork.AssociationAccessModeAudit),
		// 					HasProvisioningIssues: to.Ptr("no"),
		// 					PrivateLinkResource: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/{paasSubscriptionId}/resourceGroups/{paasResourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}"),
		// 					},
		// 					Profile: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/profiles/{profileName}"),
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.NspProvisioningStateSucceeded),
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
		// 			{
		// 				Name: to.Ptr("association3"),
		// 				Type: to.Ptr("Microsoft.Network/networkSecurityPerimeters/resourceAssociations"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/resourceAssociations/association3"),
		// 				Properties: &armnetwork.NspAssociationProperties{
		// 					AccessMode: to.Ptr(armnetwork.AssociationAccessModeLearning),
		// 					HasProvisioningIssues: to.Ptr("yes"),
		// 					PrivateLinkResource: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/{paasSubscriptionId}/resourceGroups/{paasResourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}"),
		// 					},
		// 					Profile: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/profiles/{profileName}"),
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.NspProvisioningStateFailed),
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
