package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyMetadata_List_WithTop.json
func ExamplePolicyMetadataClient_NewListPager_getCollectionOfPolicyMetadataResourcesUsingTopQueryParameter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyMetadataClient().NewListPager(&armpolicyinsights.QueryOptions{Top: to.Ptr[int32](1),
		Filter:    nil,
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    nil,
	}, nil)
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
		// page.PolicyMetadataCollection = armpolicyinsights.PolicyMetadataCollection{
		// 	Value: []*armpolicyinsights.SlimPolicyMetadata{
		// 		{
		// 			Name: to.Ptr("NIST_SP_800-53_R4_AC-2"),
		// 			Type: to.Ptr("Microsoft.PolicyInsights/policyMetadata"),
		// 			ID: to.Ptr("/providers/Microsoft.PolicyInsights/policyMetadata/NIST_SP_800-53_R4_AC-2"),
		// 			Properties: &armpolicyinsights.PolicyMetadataSlimProperties{
		// 				AdditionalContentURL: to.Ptr("https://aka.ms/NIST_SP_800-53_R4_AC-2"),
		// 				Category: to.Ptr("Access control"),
		// 				Metadata: map[string]any{
		// 				},
		// 				MetadataID: to.Ptr("NIST SP 800-53 R4 AC-2"),
		// 				Owner: to.Ptr("Shared"),
		// 				Title: to.Ptr("Account Management"),
		// 			},
		// 	}},
		// }
	}
}
