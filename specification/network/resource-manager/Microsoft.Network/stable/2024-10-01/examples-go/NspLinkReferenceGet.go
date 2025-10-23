package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a9dbb28e788355a47dc5bad3ea5f8da212b4bf6/specification/network/resource-manager/Microsoft.Network/stable/2024-10-01/examples/NspLinkReferenceGet.json
func ExampleSecurityPerimeterLinkReferencesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecurityPerimeterLinkReferencesClient().Get(ctx, "rg1", "nsp2", "link1-guid", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NspLinkReference = armnetwork.NspLinkReference{
	// 	Name: to.Ptr("link1-guid"),
	// 	Type: to.Ptr("Microsoft.Network/networkSecurityPerimeters/linkreferences"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp2/linkreferences/link1-guid"),
	// 	SystemData: &armnetwork.SecurityPerimeterSystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.344Z"); return t}()),
	// 		CreatedBy: to.Ptr("user"),
	// 		CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.344Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user"),
	// 		LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 	},
	// 	Properties: &armnetwork.NspLinkReferenceProperties{
	// 		Description: to.Ptr("Auto Approved"),
	// 		LocalInboundProfiles: []*string{
	// 			to.Ptr("*")},
	// 			LocalOutboundProfiles: []*string{
	// 				to.Ptr("*")},
	// 				ProvisioningState: to.Ptr(armnetwork.NspLinkProvisioningStateSucceeded),
	// 				RemoteInboundProfiles: []*string{
	// 					to.Ptr("*")},
	// 					RemoteOutboundProfiles: []*string{
	// 						to.Ptr("*")},
	// 						RemotePerimeterGUID: to.Ptr("guid"),
	// 						RemotePerimeterLocation: to.Ptr("westus2"),
	// 						RemotePerimeterResourceID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1"),
	// 						Status: to.Ptr(armnetwork.NspLinkStatusApproved),
	// 					},
	// 				}
}
