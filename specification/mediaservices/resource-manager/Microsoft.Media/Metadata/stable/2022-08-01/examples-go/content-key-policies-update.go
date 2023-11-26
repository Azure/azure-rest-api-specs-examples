package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-update.json
func ExampleContentKeyPoliciesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContentKeyPoliciesClient().Update(ctx, "contoso", "contosomedia", "PolicyWithClearKeyOptionAndTokenRestriction", armmediaservices.ContentKeyPolicy{
		Properties: &armmediaservices.ContentKeyPolicyProperties{
			Description: to.Ptr("Updated Policy"),
			Options: []*armmediaservices.ContentKeyPolicyOption{
				{
					Name: to.Ptr("ClearKeyOption"),
					Configuration: &armmediaservices.ContentKeyPolicyClearKeyConfiguration{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration"),
					},
					Restriction: &armmediaservices.ContentKeyPolicyOpenRestriction{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyOpenRestriction"),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ContentKeyPolicy = armmediaservices.ContentKeyPolicy{
	// 	Name: to.Ptr("PolicyWithClearKeyOptionAndTokenRestriction"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/contentKeyPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/contentKeyPolicies/PolicyWithClearKeyOptionAndTokenRestriction"),
	// 	Properties: &armmediaservices.ContentKeyPolicyProperties{
	// 		Description: to.Ptr("Updated Policy"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-01T00:00:00.000Z"); return t}()),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:30.303Z"); return t}()),
	// 		Options: []*armmediaservices.ContentKeyPolicyOption{
	// 			{
	// 				Name: to.Ptr("ClearKeyOption"),
	// 				Configuration: &armmediaservices.ContentKeyPolicyClearKeyConfiguration{
	// 					ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration"),
	// 				},
	// 				PolicyOptionID: to.Ptr("7d3f4bc1-d2bf-43a3-b02e-a7e31ab15d43"),
	// 				Restriction: &armmediaservices.ContentKeyPolicyOpenRestriction{
	// 					ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyOpenRestriction"),
	// 				},
	// 		}},
	// 		PolicyID: to.Ptr("8352435b-ebea-4681-aae7-e19277771f64"),
	// 	},
	// }
}
