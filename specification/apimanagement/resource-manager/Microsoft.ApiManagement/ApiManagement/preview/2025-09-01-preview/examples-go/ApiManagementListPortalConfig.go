package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListPortalConfig.json
func ExamplePortalConfigClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPortalConfigClient().NewListByServicePager("rg1", "apimService1", nil)
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
		// page = armapimanagement.PortalConfigClientListByServiceResponse{
		// 	PortalConfigCollection: armapimanagement.PortalConfigCollection{
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armapimanagement.PortalConfigContract{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/portalconfigs"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/portalconfigs/default"),
		// 				Properties: &armapimanagement.PortalConfigProperties{
		// 					Cors: &armapimanagement.PortalConfigCorsProperties{
		// 						AllowedOrigins: []*string{
		// 							to.Ptr("https://contoso.com"),
		// 						},
		// 					},
		// 					Csp: &armapimanagement.PortalConfigCspProperties{
		// 						AllowedSources: []*string{
		// 							to.Ptr("*.contoso.com"),
		// 						},
		// 						Mode: to.Ptr(armapimanagement.PortalSettingsCspModeReportOnly),
		// 						ReportURI: []*string{
		// 							to.Ptr("https://report.contoso.com"),
		// 						},
		// 					},
		// 					Delegation: &armapimanagement.PortalConfigDelegationProperties{
		// 						DelegateRegistration: to.Ptr(false),
		// 						DelegateSubscription: to.Ptr(false),
		// 					},
		// 					EnableBasicAuth: to.Ptr(true),
		// 					Signin: &armapimanagement.PortalConfigPropertiesSignin{
		// 						Require: to.Ptr(false),
		// 					},
		// 					Signup: &armapimanagement.PortalConfigPropertiesSignup{
		// 						TermsOfService: &armapimanagement.PortalConfigTermsOfServiceProperties{
		// 							RequireConsent: to.Ptr(false),
		// 							Text: to.Ptr("I agree to the service terms and conditions."),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
