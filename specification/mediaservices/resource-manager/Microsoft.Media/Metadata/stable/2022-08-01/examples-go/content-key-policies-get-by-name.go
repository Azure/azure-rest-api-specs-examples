package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-get-by-name.json
func ExampleContentKeyPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContentKeyPoliciesClient().Get(ctx, "contoso", "contosomedia", "PolicyWithMultipleOptions", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ContentKeyPolicy = armmediaservices.ContentKeyPolicy{
	// 	Name: to.Ptr("PolicyWithMultipleOptions"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/contentKeyPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/contentKeyPolicies/PolicyWithMultipleOptions"),
	// 	Properties: &armmediaservices.ContentKeyPolicyProperties{
	// 		Description: to.Ptr("A policy with multiple options."),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-12-01T00:00:00.000Z"); return t}()),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-02T00:00:00.000Z"); return t}()),
	// 		Options: []*armmediaservices.ContentKeyPolicyOption{
	// 			{
	// 				Configuration: &armmediaservices.ContentKeyPolicyClearKeyConfiguration{
	// 					ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration"),
	// 				},
	// 				PolicyOptionID: to.Ptr("caf1e28c-8288-4301-8c46-c0f9312c512f"),
	// 				Restriction: &armmediaservices.ContentKeyPolicyTokenRestriction{
	// 					ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyTokenRestriction"),
	// 					AlternateVerificationKeys: []armmediaservices.ContentKeyPolicyRestrictionTokenKeyClassification{
	// 					},
	// 					Audience: to.Ptr("urn:test"),
	// 					Issuer: to.Ptr("http://testacs"),
	// 					PrimaryVerificationKey: &armmediaservices.ContentKeyPolicySymmetricTokenKey{
	// 						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicySymmetricTokenKey"),
	// 						KeyValue: []byte(""),
	// 					},
	// 					RequiredClaims: []*armmediaservices.ContentKeyPolicyTokenClaim{
	// 						{
	// 							ClaimType: to.Ptr("urn:microsoft:azure:mediaservices:contentkeyidentifier"),
	// 						},
	// 						{
	// 							ClaimType: to.Ptr("DRM"),
	// 							ClaimValue: to.Ptr("Widevine"),
	// 					}},
	// 					RestrictionTokenType: to.Ptr(armmediaservices.ContentKeyPolicyRestrictionTokenTypeJwt),
	// 				},
	// 			},
	// 			{
	// 				Configuration: &armmediaservices.ContentKeyPolicyWidevineConfiguration{
	// 					ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyWidevineConfiguration"),
	// 					WidevineTemplate: to.Ptr("{\"allowed_track_types\":\"SD_HD\",\"content_key_specs\":[{\"track_type\":\"SD\",\"security_level\":1,\"required_output_protection\":{\"hdcp\":\"HDCP_V2\"}}],\"policy_overrides\":{\"can_play\":true,\"can_persist\":true,\"can_renew\":false}}"),
	// 				},
	// 				PolicyOptionID: to.Ptr("da346259-0cd6-4609-89dc-15ac131bd92f"),
	// 				Restriction: &armmediaservices.ContentKeyPolicyOpenRestriction{
	// 					ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyOpenRestriction"),
	// 				},
	// 		}},
	// 		PolicyID: to.Ptr("ed7f3d1b-cfa7-4181-b966-e0b3027eec3a"),
	// 	},
	// }
}
