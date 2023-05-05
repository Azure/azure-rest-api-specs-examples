package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fdf43f2fdacf17bd78c0621df44a5c024b61db82/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/FqdnListGlobalRulestack_List_MaximumSet_Gen.json
func ExampleFqdnListGlobalRulestackClient_NewListPager_fqdnListGlobalRulestackListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFqdnListGlobalRulestackClient().NewListPager("praval", nil)
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
		// page.FqdnListGlobalRulestackResourceListResult = armpanngfw.FqdnListGlobalRulestackResourceListResult{
		// 	Value: []*armpanngfw.FqdnListGlobalRulestackResource{
		// 		{
		// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			Type: to.Ptr("aaaaa"),
		// 			ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 			SystemData: &armpanngfw.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
		// 				CreatedBy: to.Ptr("praval"),
		// 				CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("praval"),
		// 				LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
		// 			},
		// 			Properties: &armpanngfw.FqdnObject{
		// 				Description: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 				AuditComment: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				Etag: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 				FqdnList: []*string{
		// 					to.Ptr("string1"),
		// 					to.Ptr("string2")},
		// 					ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}
