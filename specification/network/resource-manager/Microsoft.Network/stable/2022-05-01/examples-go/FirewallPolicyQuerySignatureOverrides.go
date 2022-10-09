package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/FirewallPolicyQuerySignatureOverrides.json
func ExampleFirewallPolicyIdpsSignaturesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyIdpsSignaturesClient("e747cc13-97d4-4a79-b463-42d7f4e558f2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx, "rg1", "firewallPolicy", armnetwork.IDPSQueryObject{
		Filters: []*armnetwork.FilterItems{
			{
				Field: to.Ptr("Mode"),
				Values: []*string{
					to.Ptr("Deny")},
			}},
		OrderBy: &armnetwork.OrderBy{
			Field: to.Ptr("severity"),
			Order: to.Ptr(armnetwork.FirewallPolicyIDPSQuerySortOrderAscending),
		},
		ResultsPerPage: to.Ptr[int32](20),
		Search:         to.Ptr(""),
		Skip:           to.Ptr[int32](0),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
