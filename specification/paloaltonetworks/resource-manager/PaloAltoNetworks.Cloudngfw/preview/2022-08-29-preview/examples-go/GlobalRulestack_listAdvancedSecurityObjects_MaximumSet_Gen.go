package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fdf43f2fdacf17bd78c0621df44a5c024b61db82/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/GlobalRulestack_listAdvancedSecurityObjects_MaximumSet_Gen.json
func ExampleGlobalRulestackClient_ListAdvancedSecurityObjects_globalRulestackListAdvancedSecurityObjectsMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGlobalRulestackClient().ListAdvancedSecurityObjects(ctx, "praval", armpanngfw.AdvSecurityObjectTypeEnum("globalRulestacks"), &armpanngfw.GlobalRulestackClientListAdvancedSecurityObjectsOptions{Skip: to.Ptr("a6a321"),
		Top: to.Ptr[int32](20),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AdvSecurityObjectListResponse = armpanngfw.AdvSecurityObjectListResponse{
	// 	Value: &armpanngfw.AdvSecurityObjectModel{
	// 		Type: to.Ptr("globalRulestacks"),
	// 		Entry: []*armpanngfw.NameDescriptionObject{
	// 			{
	// 				Name: to.Ptr("aaaaaaaaaa"),
	// 				Description: to.Ptr("aaaaaaaaaaaa"),
	// 		}},
	// 	},
	// }
}
