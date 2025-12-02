package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/PrefixListLocalRulestack_CreateOrUpdate_MaximumSet_Gen.json
func ExamplePrefixListLocalRulestackClient_BeginCreateOrUpdate_prefixListLocalRulestackCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("2bf4a339-294d-4c25-b0b2-ef649e9f5c27", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrefixListLocalRulestackClient().BeginCreateOrUpdate(ctx, "rgopenapi", "lrs1", "armid1", armpanngfw.PrefixListResource{
		Properties: &armpanngfw.PrefixObject{
			Description:  to.Ptr("string"),
			AuditComment: to.Ptr("comment"),
			Etag:         to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c27"),
			PrefixList: []*string{
				to.Ptr("1.0.0.0/24"),
			},
			ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpanngfw.PrefixListLocalRulestackClientCreateOrUpdateResponse{
	// 	PrefixListResource: &armpanngfw.PrefixListResource{
	// 		Name: to.Ptr("armid1"),
	// 		Type: to.Ptr("certificates"),
	// 		ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalrulestacks/armid1/certificates/armid1"),
	// 		Properties: &armpanngfw.PrefixObject{
	// 			Description: to.Ptr("string"),
	// 			AuditComment: to.Ptr("comment"),
	// 			Etag: to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c27"),
	// 			PrefixList: []*string{
	// 				to.Ptr("1.0.0.0/24"),
	// 			},
	// 			ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
	// 		},
	// 		SystemData: &armpanngfw.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 			CreatedBy: to.Ptr("praval"),
	// 			CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("praval"),
	// 			LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
