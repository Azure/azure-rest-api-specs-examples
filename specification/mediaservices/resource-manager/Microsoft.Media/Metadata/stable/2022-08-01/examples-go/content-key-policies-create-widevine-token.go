package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-create-widevine-token.json
func ExampleContentKeyPoliciesClient_CreateOrUpdate_createsAContentKeyPolicyWithWidevineOptionAndTokenRestriction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContentKeyPoliciesClient().CreateOrUpdate(ctx, "contoso", "contosomedia", "PolicyWithWidevineOptionAndJwtTokenRestriction", armmediaservices.ContentKeyPolicy{
		Properties: &armmediaservices.ContentKeyPolicyProperties{
			Description: to.Ptr("ArmPolicyDescription"),
			Options: []*armmediaservices.ContentKeyPolicyOption{
				{
					Name: to.Ptr("widevineoption"),
					Configuration: &armmediaservices.ContentKeyPolicyWidevineConfiguration{
						ODataType:        to.Ptr("#Microsoft.Media.ContentKeyPolicyWidevineConfiguration"),
						WidevineTemplate: to.Ptr("{\"allowed_track_types\":\"SD_HD\",\"content_key_specs\":[{\"track_type\":\"SD\",\"security_level\":1,\"required_output_protection\":{\"hdcp\":\"HDCP_V2\"}}],\"policy_overrides\":{\"can_play\":true,\"can_persist\":true,\"can_renew\":false}}"),
					},
					Restriction: &armmediaservices.ContentKeyPolicyTokenRestriction{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyTokenRestriction"),
						AlternateVerificationKeys: []armmediaservices.ContentKeyPolicyRestrictionTokenKeyClassification{
							&armmediaservices.ContentKeyPolicySymmetricTokenKey{
								ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicySymmetricTokenKey"),
								KeyValue:  []byte("AAAAAAAAAAAAAAAAAAAAAA=="),
							}},
						Audience: to.Ptr("urn:audience"),
						Issuer:   to.Ptr("urn:issuer"),
						PrimaryVerificationKey: &armmediaservices.ContentKeyPolicyRsaTokenKey{
							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyRsaTokenKey"),
							Exponent:  []byte("AQAB"),
							Modulus:   []byte("AQAD"),
						},
						RestrictionTokenType: to.Ptr(armmediaservices.ContentKeyPolicyRestrictionTokenTypeJwt),
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
	// 	Name: to.Ptr("PolicyWithWidevineOptionAndJwtTokenRestriction"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/contentKeyPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/contentKeyPolicies/PolicyWithWidevineOptionAndJwtTokenRestriction"),
	// 	Properties: &armmediaservices.ContentKeyPolicyProperties{
	// 		Description: to.Ptr("ArmPolicyDescription"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:29.663Z"); return t}()),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:29.663Z"); return t}()),
	// 		Options: []*armmediaservices.ContentKeyPolicyOption{
	// 			{
	// 				Name: to.Ptr("widevineoption"),
	// 				Configuration: &armmediaservices.ContentKeyPolicyWidevineConfiguration{
	// 					ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyWidevineConfiguration"),
	// 					WidevineTemplate: to.Ptr("{\"allowed_track_types\":\"SD_HD\",\"content_key_specs\":[{\"track_type\":\"SD\",\"security_level\":1,\"required_output_protection\":{\"hdcp\":\"HDCP_V2\"}}],\"policy_overrides\":{\"can_play\":true,\"can_persist\":true,\"can_renew\":false}}"),
	// 				},
	// 				PolicyOptionID: to.Ptr("26fee004-8dfa-4828-bcad-5e63c637534f"),
	// 				Restriction: &armmediaservices.ContentKeyPolicyTokenRestriction{
	// 					ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyTokenRestriction"),
	// 					AlternateVerificationKeys: []armmediaservices.ContentKeyPolicyRestrictionTokenKeyClassification{
	// 						&armmediaservices.ContentKeyPolicySymmetricTokenKey{
	// 							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicySymmetricTokenKey"),
	// 							KeyValue: []byte(""),
	// 					}},
	// 					Audience: to.Ptr("urn:audience"),
	// 					Issuer: to.Ptr("urn:issuer"),
	// 					PrimaryVerificationKey: &armmediaservices.ContentKeyPolicyRsaTokenKey{
	// 						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyRsaTokenKey"),
	// 						Exponent: []byte(""),
	// 						Modulus: []byte(""),
	// 					},
	// 					RequiredClaims: []*armmediaservices.ContentKeyPolicyTokenClaim{
	// 					},
	// 					RestrictionTokenType: to.Ptr(armmediaservices.ContentKeyPolicyRestrictionTokenTypeJwt),
	// 				},
	// 		}},
	// 		PolicyID: to.Ptr("bad1d030-7d5c-4643-8f1e-49807a4bf64c"),
	// 	},
	// }
}
