package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/CertificateObjectLocalRulestack_CreateOrUpdate_MinimumSet_Gen.json
func ExampleCertificateObjectLocalRulestackClient_BeginCreateOrUpdate_certificateObjectLocalRulestackCreateOrUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCertificateObjectLocalRulestackClient().BeginCreateOrUpdate(ctx, "rgopenapi", "lrs1", "armid1", armpanngfw.CertificateObjectLocalRulestackResource{
		Properties: &armpanngfw.CertificateObject{
			CertificateSelfSigned: to.Ptr(armpanngfw.BooleanEnumTRUE),
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
	// res.CertificateObjectLocalRulestackResource = armpanngfw.CertificateObjectLocalRulestackResource{
	// 	ID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/acctest4136/providers/PaloAltoNetworks.Cloudngfw/localrulestacks/lrs1/certificates/armid1"),
	// 	Properties: &armpanngfw.CertificateObject{
	// 		CertificateSelfSigned: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 	},
	// }
}
