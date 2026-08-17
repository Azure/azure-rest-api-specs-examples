package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/ExpressRouteLagMemberGet.json
func ExampleExpressRouteLagsClient_MembersGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRouteLagsClient().MembersGet(ctx, "rg1", "lagName", "linkName", "memberName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ExpressRouteLagsClientMembersGetResponse{
	// 	ExpressRouteLagMember: armnetwork.ExpressRouteLagMember{
	// 		Name: to.Ptr("memberName"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName/links/linkName/members/memberName"),
	// 		Etag: to.Ptr("W/\"f6c63a53-9734-4731-8fd1-ef3ca14bd6b6\""),
	// 		Properties: &armnetwork.ExpressRouteLagMemberPropertiesFormat{
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			InterfaceName: to.Ptr("Ethernet0/0"),
	// 			PatchPanelID: to.Ptr("patchPanelId1"),
	// 			RackID: to.Ptr("rackId1"),
	// 			ColoLocation: to.Ptr(""),
	// 			ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 			AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
	// 		},
	// 		Type: to.Ptr("Microsoft.Network/expressRouteLags/links/members"),
	// 	},
	// }
}
