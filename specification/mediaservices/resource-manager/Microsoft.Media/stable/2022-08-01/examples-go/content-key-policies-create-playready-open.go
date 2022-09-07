package armmediaservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/content-key-policies-create-playready-open.json
func ExampleContentKeyPoliciesClient_CreateOrUpdate_createsAContentKeyPolicyWithPlayReadyOptionAndOpenRestriction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewContentKeyPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "contoso", "contosomedia", "PolicyWithPlayReadyOptionAndOpenRestriction", armmediaservices.ContentKeyPolicy{
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
								BeginDate:        to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-16T18:22:53.46Z"); return t }()),
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
	// TODO: use response item
	_ = res
}
