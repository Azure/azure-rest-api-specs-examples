package armartifactsigning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/artifactsigning/armartifactsigning"
)

// Generated from example definition: 2025-10-13/CertificateProfiles_Get.json
func ExampleCertificateProfilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armartifactsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
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
	// res = armartifactsigning.CertificateProfilesClientGetResponse{
	// 	CertificateProfile: &armartifactsigning.CertificateProfile{
	// 		Name: to.Ptr("profileA"),
	// 		Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts/certificateProfiles"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/MyAccount/certificateProfiles/profileA"),
	// 		Properties: &armartifactsigning.CertificateProfileProperties{
	// 			Certificates: []*armartifactsigning.Certificate{
	// 				{
	// 					CreatedDate: to.Ptr("3/14/2023 5:27:49 PM"),
	// 					ExpiryDate: to.Ptr("3/17/2023 5:27:49 PM"),
	// 					EnhancedKeyUsage: to.Ptr("1.3.6.1.4.1.311.yy.xxxxxxxx.xxxxxxxx.xxxxxxxxx.xxxxxxxx"),
	// 					SerialNumber: to.Ptr("xxxxxxxxxxxxxxxxxx"),
	// 					Status: to.Ptr(armartifactsigning.CertificateStatusActive),
	// 					SubjectName: to.Ptr("CN=Contoso Inc, O=Contoso Inc, L=New York, S=New York, C=US"),
	// 					Thumbprint: to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
	// 				},
	// 			},
	// 			IdentityValidationID: to.Ptr("123456"),
	// 			IncludeCity: to.Ptr(false),
	// 			IncludeCountry: to.Ptr(false),
	// 			IncludePostalCode: to.Ptr(true),
	// 			IncludeState: to.Ptr(false),
	// 			IncludeStreetAddress: to.Ptr(false),
	// 			ProfileType: to.Ptr(armartifactsigning.ProfileTypePublicTrust),
	// 			ProvisioningState: to.Ptr(armartifactsigning.ProvisioningStateSucceeded),
	// 			Status: to.Ptr(armartifactsigning.CertificateProfileStatusActive),
	// 		},
	// 	},
	// }
}
