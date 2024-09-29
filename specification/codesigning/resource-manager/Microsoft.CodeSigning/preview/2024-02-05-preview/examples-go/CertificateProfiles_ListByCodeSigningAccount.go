package armtrustedsigning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trustedsigning/armtrustedsigning"
)

// Generated from example definition: 2024-02-05-preview/CertificateProfiles_ListByCodeSigningAccount.json
func ExampleCertificateProfilesClient_NewListByCodeSigningAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrustedsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCertificateProfilesClient().NewListByCodeSigningAccountPager("MyResourceGroup", "MyAccount", nil)
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
		// page = armtrustedsigning.CertificateProfilesClientListByCodeSigningAccountResponse{
		// 	CertificateProfileListResult: armtrustedsigning.CertificateProfileListResult{
		// 		Value: []*armtrustedsigning.CertificateProfile{
		// 			{
		// 				Name: to.Ptr("profileA"),
		// 				Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts/certificateProfiles"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/profileA"),
		// 				Properties: &armtrustedsigning.CertificateProfileProperties{
		// 					Certificates: []*armtrustedsigning.Certificate{
		// 						{
		// 							CreatedDate: to.Ptr("3/14/2023 5:27:49 PM"),
		// 							ExpiryDate: to.Ptr("3/17/2023 5:27:49 PM"),
		// 							SerialNumber: to.Ptr("xxxxxxxxxxxxxxxxxx"),
		// 							Status: to.Ptr(armtrustedsigning.CertificateStatusActive),
		// 							SubjectName: to.Ptr("CN=Contoso Inc, O=Contoso Inc, L=New York, S=New York, C=US"),
		// 							Thumbprint: to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		// 						},
		// 					},
		// 					City: to.Ptr("Dallas"),
		// 					CommonName: to.Ptr("Microsoft Corporation"),
		// 					Country: to.Ptr("US"),
		// 					EnhancedKeyUsage: to.Ptr("1.3.6.1.4.1.311.yy.xxxxxxxx.xxxxxxxx.xxxxxxxxx.xxxxxxxx"),
		// 					IdentityValidationID: to.Ptr("123456"),
		// 					IncludeCity: to.Ptr(false),
		// 					IncludeCountry: to.Ptr(false),
		// 					IncludePostalCode: to.Ptr(true),
		// 					IncludeState: to.Ptr(false),
		// 					IncludeStreetAddress: to.Ptr(false),
		// 					Organization: to.Ptr("Microsoft Corporation"),
		// 					PostalCode: to.Ptr("560090"),
		// 					ProfileType: to.Ptr(armtrustedsigning.ProfileTypePublicTrust),
		// 					ProvisioningState: to.Ptr(armtrustedsigning.ProvisioningStateSucceeded),
		// 					State: to.Ptr("Texas"),
		// 					Status: to.Ptr(armtrustedsigning.CertificateProfileStatusActive),
		// 					StreetAddress: to.Ptr("123 Bluebonnet"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
