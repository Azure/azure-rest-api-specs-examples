package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/CertificateObjectGlobalRulestack_CreateOrUpdate_MaximumSet_Gen.json
func ExampleCertificateObjectGlobalRulestackClient_BeginCreateOrUpdate_certificateObjectGlobalRulestackCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCertificateObjectGlobalRulestackClient().BeginCreateOrUpdate(ctx, "praval", "armid1", armpanngfw.CertificateObjectGlobalRulestackResource{
		Properties: &armpanngfw.CertificateObject{
			Description:                 to.Ptr("description"),
			AuditComment:                to.Ptr("comment"),
			CertificateSelfSigned:       to.Ptr(armpanngfw.BooleanEnumTRUE),
			CertificateSignerResourceID: to.Ptr(""),
			Etag:                        to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c27"),
			ProvisioningState:           to.Ptr(armpanngfw.ProvisioningStateAccepted),
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
	// res.CertificateObjectGlobalRulestackResource = armpanngfw.CertificateObjectGlobalRulestackResource{
	// 	Name: to.Ptr("armid1"),
	// 	Type: to.Ptr("certificates"),
	// 	ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalrulestacks/armid1/certificates/armid1"),
	// 	SystemData: &armpanngfw.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		CreatedBy: to.Ptr("praval"),
	// 		CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("praval"),
	// 		LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 	},
	// 	Properties: &armpanngfw.CertificateObject{
	// 		Description: to.Ptr("description"),
	// 		AuditComment: to.Ptr("comment"),
	// 		CertificateSelfSigned: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 		CertificateSignerResourceID: to.Ptr(""),
	// 		Etag: to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c27"),
	// 		ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateSucceeded),
	// 	},
	// }
}
