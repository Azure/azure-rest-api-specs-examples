package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-list-all.json
func ExampleContentKeyPoliciesClient_NewListPager_listsAllContentKeyPolicies() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContentKeyPoliciesClient().NewListPager("contoso", "contosomedia", &armmediaservices.ContentKeyPoliciesClientListOptions{Filter: nil,
		Top:     nil,
		Orderby: nil,
	})
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
		// page.ContentKeyPolicyCollection = armmediaservices.ContentKeyPolicyCollection{
		// 	Value: []*armmediaservices.ContentKeyPolicy{
		// 		{
		// 			Name: to.Ptr("PolicyWithClearKeyOptionAndTokenRestriction"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/contentKeyPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/contentKeyPolicies/PolicyWithClearKeyOptionAndTokenRestriction"),
		// 			Properties: &armmediaservices.ContentKeyPolicyProperties{
		// 				Description: to.Ptr("A policy with one ClearKey option and Open Restriction."),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-01T00:00:00.000Z"); return t}()),
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-01T00:00:00.000Z"); return t}()),
		// 				Options: []*armmediaservices.ContentKeyPolicyOption{
		// 					{
		// 						Configuration: &armmediaservices.ContentKeyPolicyClearKeyConfiguration{
		// 							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration"),
		// 						},
		// 						PolicyOptionID: to.Ptr("a3448d09-567a-4642-8309-d17e846be59f"),
		// 						Restriction: &armmediaservices.ContentKeyPolicyTokenRestriction{
		// 							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyTokenRestriction"),
		// 							AlternateVerificationKeys: []armmediaservices.ContentKeyPolicyRestrictionTokenKeyClassification{
		// 							},
		// 							Audience: to.Ptr("urn:test"),
		// 							Issuer: to.Ptr("http://testacs"),
		// 							PrimaryVerificationKey: &armmediaservices.ContentKeyPolicySymmetricTokenKey{
		// 								ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicySymmetricTokenKey"),
		// 								KeyValue: []byte(""),
		// 							},
		// 							RequiredClaims: []*armmediaservices.ContentKeyPolicyTokenClaim{
		// 								{
		// 									ClaimType: to.Ptr("urn:microsoft:azure:mediaservices:contentkeyidentifier"),
		// 								},
		// 								{
		// 									ClaimType: to.Ptr("DRM"),
		// 									ClaimValue: to.Ptr("Widevine"),
		// 							}},
		// 							RestrictionTokenType: to.Ptr(armmediaservices.ContentKeyPolicyRestrictionTokenTypeJwt),
		// 						},
		// 				}},
		// 				PolicyID: to.Ptr("8352435b-ebea-4681-aae7-e19277771f64"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PolicyWithMultipleOptions"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/contentKeyPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/contentKeyPolicies/PolicyWithMultipleOptions"),
		// 			Properties: &armmediaservices.ContentKeyPolicyProperties{
		// 				Description: to.Ptr("A policy with multiple options."),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-12-01T00:00:00.000Z"); return t}()),
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-02T00:00:00.000Z"); return t}()),
		// 				Options: []*armmediaservices.ContentKeyPolicyOption{
		// 					{
		// 						Configuration: &armmediaservices.ContentKeyPolicyClearKeyConfiguration{
		// 							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration"),
		// 						},
		// 						PolicyOptionID: to.Ptr("caf1e28c-8288-4301-8c46-c0f9312c512f"),
		// 						Restriction: &armmediaservices.ContentKeyPolicyTokenRestriction{
		// 							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyTokenRestriction"),
		// 							AlternateVerificationKeys: []armmediaservices.ContentKeyPolicyRestrictionTokenKeyClassification{
		// 							},
		// 							Audience: to.Ptr("urn:test"),
		// 							Issuer: to.Ptr("http://testacs"),
		// 							PrimaryVerificationKey: &armmediaservices.ContentKeyPolicySymmetricTokenKey{
		// 								ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicySymmetricTokenKey"),
		// 								KeyValue: []byte(""),
		// 							},
		// 							RequiredClaims: []*armmediaservices.ContentKeyPolicyTokenClaim{
		// 								{
		// 									ClaimType: to.Ptr("urn:microsoft:azure:mediaservices:contentkeyidentifier"),
		// 								},
		// 								{
		// 									ClaimType: to.Ptr("DRM"),
		// 									ClaimValue: to.Ptr("Widevine"),
		// 							}},
		// 							RestrictionTokenType: to.Ptr(armmediaservices.ContentKeyPolicyRestrictionTokenTypeJwt),
		// 						},
		// 					},
		// 					{
		// 						Configuration: &armmediaservices.ContentKeyPolicyWidevineConfiguration{
		// 							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyWidevineConfiguration"),
		// 							WidevineTemplate: to.Ptr("{\"allowed_track_types\":\"SD_HD\",\"content_key_specs\":[{\"track_type\":\"SD\",\"security_level\":1,\"required_output_protection\":{\"hdcp\":\"HDCP_V2\"}}],\"policy_overrides\":{\"can_play\":true,\"can_persist\":true,\"can_renew\":false}}"),
		// 						},
		// 						PolicyOptionID: to.Ptr("da346259-0cd6-4609-89dc-15ac131bd92f"),
		// 						Restriction: &armmediaservices.ContentKeyPolicyOpenRestriction{
		// 							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyOpenRestriction"),
		// 						},
		// 				}},
		// 				PolicyID: to.Ptr("ed7f3d1b-cfa7-4181-b966-e0b3027eec3a"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PolicyWithPlayReadyOptionAndOpenRestriction"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/contentKeyPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/contentKeyPolicies/PolicyWithPlayReadyOptionAndOpenRestriction"),
		// 			Properties: &armmediaservices.ContentKeyPolicyProperties{
		// 				Description: to.Ptr("A policy with one PlayReady option and Open Restriction."),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2012-11-01T00:00:00.000Z"); return t}()),
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2012-11-01T00:00:00.000Z"); return t}()),
		// 				Options: []*armmediaservices.ContentKeyPolicyOption{
		// 					{
		// 						Configuration: &armmediaservices.ContentKeyPolicyPlayReadyConfiguration{
		// 							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyPlayReadyConfiguration"),
		// 							Licenses: []*armmediaservices.ContentKeyPolicyPlayReadyLicense{
		// 								{
		// 									AllowTestDevices: to.Ptr(false),
		// 									ContentKeyLocation: &armmediaservices.ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader{
		// 										ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader"),
		// 									},
		// 									ContentType: to.Ptr(armmediaservices.ContentKeyPolicyPlayReadyContentTypeUnspecified),
		// 									LicenseType: to.Ptr(armmediaservices.ContentKeyPolicyPlayReadyLicenseTypeNonPersistent),
		// 									PlayRight: &armmediaservices.ContentKeyPolicyPlayReadyPlayRight{
		// 										AllowPassingVideoContentToUnknownOutput: to.Ptr(armmediaservices.ContentKeyPolicyPlayReadyUnknownOutputPassingOptionNotAllowed),
		// 										DigitalVideoOnlyContentRestriction: to.Ptr(false),
		// 										ImageConstraintForAnalogComponentVideoRestriction: to.Ptr(false),
		// 										ImageConstraintForAnalogComputerMonitorRestriction: to.Ptr(false),
		// 									},
		// 									SecurityLevel: to.Ptr(armmediaservices.SecurityLevelSL2000),
		// 							}},
		// 							ResponseCustomData: to.Ptr("testCustomData"),
		// 						},
		// 						PolicyOptionID: to.Ptr("294a833f-f128-48be-9edf-8d1bb5b35ff3"),
		// 						Restriction: &armmediaservices.ContentKeyPolicyOpenRestriction{
		// 							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyOpenRestriction"),
		// 						},
		// 				}},
		// 				PolicyID: to.Ptr("a9bacd1d-60f5-4af3-8d2b-cf46ca5c9b04"),
		// 			},
		// 	}},
		// }
	}
}
