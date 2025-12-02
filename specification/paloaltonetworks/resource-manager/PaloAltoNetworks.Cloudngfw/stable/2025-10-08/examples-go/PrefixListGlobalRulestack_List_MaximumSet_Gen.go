package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/PrefixListGlobalRulestack_List_MaximumSet_Gen.json
func ExamplePrefixListGlobalRulestackClient_NewListPager_prefixListGlobalRulestackListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrefixListGlobalRulestackClient().NewListPager("praval", nil)
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
		// page = armpanngfw.PrefixListGlobalRulestackClientListResponse{
		// 	PrefixListGlobalRulestackResourceListResult: armpanngfw.PrefixListGlobalRulestackResourceListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/providers/PaloAltoNetworks.Cloudngfw/globalrulestacks/praval/prefixlists?api-version=2025-10-08&$skiptoken=xyz"),
		// 		Value: []*armpanngfw.PrefixListGlobalRulestackResource{
		// 			{
		// 				Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				ID: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 				Properties: &armpanngfw.PrefixObject{
		// 					Description: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 					AuditComment: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					Etag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 					PrefixList: []*string{
		// 						to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					},
		// 					ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
		// 				},
		// 				SystemData: &armpanngfw.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
		// 					CreatedBy: to.Ptr("praval"),
		// 					CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("praval"),
		// 					LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
