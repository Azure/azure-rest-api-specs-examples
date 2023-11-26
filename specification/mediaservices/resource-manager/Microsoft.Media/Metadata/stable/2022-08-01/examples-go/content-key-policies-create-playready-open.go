package armmediaservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-create-playready-open.json
func ExampleContentKeyPoliciesClient_CreateOrUpdate_createsAContentKeyPolicyWithPlayReadyOptionAndOpenRestriction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContentKeyPoliciesClient().CreateOrUpdate(ctx, "contoso", "contosomedia", "PolicyWithPlayReadyOptionAndOpenRestriction", armmediaservices.ContentKeyPolicy{
		Properties: &armmediaservices.ContentKeyPolicyProperties{
			Description: to.Ptr("ArmPolicyDescription"),
			Options: []*armmediaservices.ContentKeyPolicyOption{
				{
					Name: to.Ptr("ArmPolicyOptionName"),
					Configuration: &armmediaservices.ContentKeyPolicyPlayReadyConfiguration{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyPlayReadyConfiguration"),
						Licenses: []*armmediaservices.ContentKeyPolicyPlayReadyLicense{
							{
								AllowTestDevices: to.Ptr(true),
								BeginDate:        to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-16T18:22:53.460Z"); return t }()),
								ContentKeyLocation: &armmediaservices.ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader{
									ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader"),
								},
								ContentType: to.Ptr(armmediaservices.ContentKeyPolicyPlayReadyContentTypeUltraVioletDownload),
								LicenseType: to.Ptr(armmediaservices.ContentKeyPolicyPlayReadyLicenseTypePersistent),
								PlayRight: &armmediaservices.ContentKeyPolicyPlayReadyPlayRight{
									AllowPassingVideoContentToUnknownOutput:            to.Ptr(armmediaservices.ContentKeyPolicyPlayReadyUnknownOutputPassingOptionNotAllowed),
									DigitalVideoOnlyContentRestriction:                 to.Ptr(false),
									ImageConstraintForAnalogComponentVideoRestriction:  to.Ptr(true),
									ImageConstraintForAnalogComputerMonitorRestriction: to.Ptr(false),
									ScmsRestriction: to.Ptr[int32](2),
								},
								SecurityLevel: to.Ptr(armmediaservices.SecurityLevelSL150),
							}},
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
	// 	Name: to.Ptr("PolicyWithPlayReadyOptionAndOpenRestriction"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/contentKeyPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/contentKeyPolicies/PolicyWithPlayReadyOptionAndOpenRestriction"),
	// 	Properties: &armmediaservices.ContentKeyPolicyProperties{
	// 		Description: to.Ptr("ArmPolicyDescription"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2012-11-01T00:00:00.000Z"); return t}()),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:29.510Z"); return t}()),
	// 		Options: []*armmediaservices.ContentKeyPolicyOption{
	// 			{
	// 				Name: to.Ptr("ArmPolicyOptionName"),
	// 				Configuration: &armmediaservices.ContentKeyPolicyPlayReadyConfiguration{
	// 					ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyPlayReadyConfiguration"),
	// 					Licenses: []*armmediaservices.ContentKeyPolicyPlayReadyLicense{
	// 						{
	// 							AllowTestDevices: to.Ptr(true),
	// 							BeginDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-16T18:22:53.460Z"); return t}()),
	// 							ContentKeyLocation: &armmediaservices.ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader{
	// 								ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader"),
	// 							},
	// 							ContentType: to.Ptr(armmediaservices.ContentKeyPolicyPlayReadyContentTypeUltraVioletDownload),
	// 							LicenseType: to.Ptr(armmediaservices.ContentKeyPolicyPlayReadyLicenseTypePersistent),
	// 							PlayRight: &armmediaservices.ContentKeyPolicyPlayReadyPlayRight{
	// 								AllowPassingVideoContentToUnknownOutput: to.Ptr(armmediaservices.ContentKeyPolicyPlayReadyUnknownOutputPassingOptionNotAllowed),
	// 								DigitalVideoOnlyContentRestriction: to.Ptr(false),
	// 								ImageConstraintForAnalogComponentVideoRestriction: to.Ptr(true),
	// 								ImageConstraintForAnalogComputerMonitorRestriction: to.Ptr(false),
	// 								ScmsRestriction: to.Ptr[int32](2),
	// 							},
	// 							SecurityLevel: to.Ptr(armmediaservices.SecurityLevelSL150),
	// 					}},
	// 				},
	// 				PolicyOptionID: to.Ptr("c52f9af0-1f53-4775-8edb-af2d9a6e28cd"),
	// 				Restriction: &armmediaservices.ContentKeyPolicyOpenRestriction{
	// 					ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyOpenRestriction"),
	// 				},
	// 		}},
	// 		PolicyID: to.Ptr("a9bacd1d-60f5-4af3-8d2b-cf46ca5c9b04"),
	// 	},
	// }
}
