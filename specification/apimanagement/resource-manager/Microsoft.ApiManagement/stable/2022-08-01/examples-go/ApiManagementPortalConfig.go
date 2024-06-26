package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementPortalConfig.json
func ExamplePortalConfigClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPortalConfigClient().Get(ctx, "rg1", "apimService1", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PortalConfigContract = armapimanagement.PortalConfigContract{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/portalconfigs"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/portalconfigs/default"),
	// 	Properties: &armapimanagement.PortalConfigProperties{
	// 		Cors: &armapimanagement.PortalConfigCorsProperties{
	// 			AllowedOrigins: []*string{
	// 				to.Ptr("https://contoso.com")},
	// 			},
	// 			Csp: &armapimanagement.PortalConfigCspProperties{
	// 				AllowedSources: []*string{
	// 					to.Ptr("*.contoso.com")},
	// 					Mode: to.Ptr(armapimanagement.PortalSettingsCspModeReportOnly),
	// 					ReportURI: []*string{
	// 						to.Ptr("https://report.contoso.com")},
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
	// 			}
}
