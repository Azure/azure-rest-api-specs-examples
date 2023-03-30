package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyMetadata_GetResource.json
func ExamplePolicyMetadataClient_GetResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyMetadataClient().GetResource(ctx, "NIST_SP_800-53_R4_AC-2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PolicyMetadata = armpolicyinsights.PolicyMetadata{
	// 	Name: to.Ptr("NIST_SP_800-53_R4_AC-2"),
	// 	Type: to.Ptr("Microsoft.PolicyInsights/policyMetadata"),
	// 	ID: to.Ptr("/providers/Microsoft.PolicyInsights/policyMetadata/NIST_SP_800-53_R4_AC-2"),
	// 	Properties: &armpolicyinsights.PolicyMetadataProperties{
	// 		AdditionalContentURL: to.Ptr("https://aka.ms/NIST_SP_800-53_R4_AC-2"),
	// 		Category: to.Ptr("Access control"),
	// 		Metadata: map[string]any{
	// 		},
	// 		MetadataID: to.Ptr("NIST SP 800-53 R4 AC-2"),
	// 		Owner: to.Ptr("Shared"),
	// 		Title: to.Ptr("Account Management"),
	// 		Description: to.Ptr("Description of NIST SP 800-53 R4 AC-2"),
	// 		Requirements: to.Ptr("List the requirements for NIST SP 800-53 R4 AC-2"),
	// 	},
	// }
}
