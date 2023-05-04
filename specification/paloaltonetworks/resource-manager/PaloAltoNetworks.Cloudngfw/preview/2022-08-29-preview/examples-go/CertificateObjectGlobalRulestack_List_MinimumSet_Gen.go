package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fdf43f2fdacf17bd78c0621df44a5c024b61db82/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/CertificateObjectGlobalRulestack_List_MinimumSet_Gen.json
func ExampleCertificateObjectGlobalRulestackClient_NewListPager_certificateObjectGlobalRulestackListMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCertificateObjectGlobalRulestackClient().NewListPager("praval", nil)
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
		// page.CertificateObjectGlobalRulestackResourceListResult = armpanngfw.CertificateObjectGlobalRulestackResourceListResult{
		// 	Value: []*armpanngfw.CertificateObjectGlobalRulestackResource{
		// 		{
		// 			ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/praval/certificates/certificates1"),
		// 			Properties: &armpanngfw.CertificateObject{
		// 				CertificateSelfSigned: to.Ptr(armpanngfw.BooleanEnumTRUE),
		// 			},
		// 	}},
		// }
	}
}
