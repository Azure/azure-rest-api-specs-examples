package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f41d0c9332078cb2ef07b749081d94915255ada5/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/privateLinkScope/PrivateLinkScopes_GetValidation.json
func ExamplePrivateLinkScopesClient_GetValidationDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkScopesClient().GetValidationDetails(ctx, "wus2", "f5dc51d3-92ed-4d7e-947a-775ea79b4919", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkScopeValidationDetails = armhybridcompute.PrivateLinkScopeValidationDetails{
	// 	ConnectionDetails: []*armhybridcompute.ConnectionDetail{
	// 		{
	// 			GroupID: to.Ptr("groupId"),
	// 			ID: to.Ptr("id"),
	// 			LinkIdentifier: to.Ptr("linkId"),
	// 			MemberName: to.Ptr("memberName"),
	// 			PrivateIPAddress: to.Ptr("ip"),
	// 	}},
	// 	ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.hybridcompute/privateLinkScopes/my-privatelinkscope"),
	// 	PublicNetworkAccess: to.Ptr(armhybridcompute.PublicNetworkAccessTypeDisabled),
	// }
}
