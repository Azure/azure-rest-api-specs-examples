package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-create-nodrm-token.json
func ExampleContentKeyPoliciesClient_CreateOrUpdate_createsAContentKeyPolicyWithClearKeyOptionAndTokenRestriction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContentKeyPoliciesClient().CreateOrUpdate(ctx, "contoso", "contosomedia", "PolicyWithClearKeyOptionAndSwtTokenRestriction", armmediaservices.ContentKeyPolicy{
		Properties: &armmediaservices.ContentKeyPolicyProperties{
			Description: to.Ptr("ArmPolicyDescription"),
			Options: []*armmediaservices.ContentKeyPolicyOption{
				{
					Name: to.Ptr("ClearKeyOption"),
					Configuration: &armmediaservices.ContentKeyPolicyClearKeyConfiguration{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration"),
					},
					Restriction: &armmediaservices.ContentKeyPolicyTokenRestriction{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyTokenRestriction"),
						Audience:  to.Ptr("urn:audience"),
						Issuer:    to.Ptr("urn:issuer"),
						PrimaryVerificationKey: &armmediaservices.ContentKeyPolicySymmetricTokenKey{
							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicySymmetricTokenKey"),
							KeyValue:  []byte("AAAAAAAAAAAAAAAAAAAAAA=="),
						},
						RestrictionTokenType: to.Ptr(armmediaservices.ContentKeyPolicyRestrictionTokenTypeSwt),
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
	// 	Name: to.Ptr("PolicyWithClearKeyOptionAndSwtTokenRestriction"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/contentKeyPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/contentKeyPolicies/PolicyWithClearKeyOptionAndSwtTokenRestriction"),
	// 	Properties: &armmediaservices.ContentKeyPolicyProperties{
	// 		Description: to.Ptr("ArmPolicyDescription"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:29.837Z"); return t}()),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:29.837Z"); return t}()),
	// 		Options: []*armmediaservices.ContentKeyPolicyOption{
	// 			{
	// 				Name: to.Ptr("ClearKeyOption"),
	// 				Configuration: &armmediaservices.ContentKeyPolicyClearKeyConfiguration{
	// 					ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration"),
	// 				},
	// 				PolicyOptionID: to.Ptr("e7d4d465-b6f7-4830-9a21-74a7326ef797"),
	// 				Restriction: &armmediaservices.ContentKeyPolicyTokenRestriction{
	// 					ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyTokenRestriction"),
	// 					AlternateVerificationKeys: []armmediaservices.ContentKeyPolicyRestrictionTokenKeyClassification{
	// 					},
	// 					Audience: to.Ptr("urn:audience"),
	// 					Issuer: to.Ptr("urn:issuer"),
	// 					PrimaryVerificationKey: &armmediaservices.ContentKeyPolicySymmetricTokenKey{
	// 						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicySymmetricTokenKey"),
	// 						KeyValue: []byte(""),
	// 					},
	// 					RequiredClaims: []*armmediaservices.ContentKeyPolicyTokenClaim{
	// 					},
	// 					RestrictionTokenType: to.Ptr(armmediaservices.ContentKeyPolicyRestrictionTokenTypeSwt),
	// 				},
	// 		}},
	// 		PolicyID: to.Ptr("2926c1bc-4dec-4a11-9d19-3f99006530a9"),
	// 	},
	// }
}
