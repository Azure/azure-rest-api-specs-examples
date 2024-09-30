package armtrustedsigning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trustedsigning/armtrustedsigning"
)

// Generated from example definition: 2024-02-05-preview/CertificateProfiles_Get.json
func ExampleCertificateProfilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrustedsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificateProfilesClient().Get(ctx, "MyResourceGroup", "MyAccount", "profileA", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armtrustedsigning.CertificateProfilesClientGetResponse{
	// 	CertificateProfile: &armtrustedsigning.CertificateProfile{
	// 		Name: to.Ptr("profileA"),
	// 		Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts/certificateProfiles"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/MyAccount/certificateProfiles/profileA"),
	// 		Properties: &armtrustedsigning.CertificateProfileProperties{
	// 			Certificates: []*armtrustedsigning.Certificate{
	// 				{
	// 					CreatedDate: to.Ptr("3/14/2023 5:27:49 PM"),
	// 					ExpiryDate: to.Ptr("3/17/2023 5:27:49 PM"),
	// 					SerialNumber: to.Ptr("xxxxxxxxxxxxxxxxxx"),
	// 					Status: to.Ptr(armtrustedsigning.CertificateStatusActive),
	// 					SubjectName: to.Ptr("CN=Contoso Inc, O=Contoso Inc, L=New York, S=New York, C=US"),
	// 					Thumbprint: to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
	// 				},
	// 			},
	// 			City: to.Ptr("Dallas"),
	// 			CommonName: to.Ptr("Contoso Inc"),
	// 			Country: to.Ptr("US"),
	// 			EnhancedKeyUsage: to.Ptr("1.3.6.1.4.1.311.yy.xxxxxxxx.xxxxxxxx.xxxxxxxxx.xxxxxxxx"),
	// 			IdentityValidationID: to.Ptr("123456"),
	// 			IncludeCity: to.Ptr(false),
	// 			IncludeCountry: to.Ptr(false),
	// 			IncludePostalCode: to.Ptr(true),
	// 			IncludeState: to.Ptr(false),
	// 			IncludeStreetAddress: to.Ptr(false),
	// 			Organization: to.Ptr("Contoso Inc"),
	// 			PostalCode: to.Ptr("560090"),
	// 			ProfileType: to.Ptr(armtrustedsigning.ProfileTypePublicTrust),
	// 			ProvisioningState: to.Ptr(armtrustedsigning.ProvisioningStateSucceeded),
	// 			State: to.Ptr("Texas"),
	// 			Status: to.Ptr(armtrustedsigning.CertificateProfileStatusActive),
	// 			StreetAddress: to.Ptr("123 Bluebonnet"),
	// 		},
	// 	},
	// }
}
