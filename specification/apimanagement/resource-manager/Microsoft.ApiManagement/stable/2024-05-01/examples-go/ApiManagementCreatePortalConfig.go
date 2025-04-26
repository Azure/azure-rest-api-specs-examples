package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreatePortalConfig.json
func ExamplePortalConfigClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPortalConfigClient().CreateOrUpdate(ctx, "rg1", "apimService1", "default", "*", armapimanagement.PortalConfigContract{
		Properties: &armapimanagement.PortalConfigProperties{
			Cors: &armapimanagement.PortalConfigCorsProperties{
				AllowedOrigins: []*string{
					to.Ptr("https://contoso.com")},
			},
			Csp: &armapimanagement.PortalConfigCspProperties{
				AllowedSources: []*string{
					to.Ptr("*.contoso.com")},
				Mode: to.Ptr(armapimanagement.PortalSettingsCspModeReportOnly),
				ReportURI: []*string{
					to.Ptr("https://report.contoso.com")},
			},
			Delegation: &armapimanagement.PortalConfigDelegationProperties{
				DelegateRegistration: to.Ptr(false),
				DelegateSubscription: to.Ptr(false),
			},
			EnableBasicAuth: to.Ptr(true),
			Signin: &armapimanagement.PortalConfigPropertiesSignin{
				Require: to.Ptr(false),
			},
			Signup: &armapimanagement.PortalConfigPropertiesSignup{
				TermsOfService: &armapimanagement.PortalConfigTermsOfServiceProperties{
					RequireConsent: to.Ptr(false),
					Text:           to.Ptr("I agree to the service terms and conditions."),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PortalConfigContract = armapimanagement.PortalConfigContract{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/portalconfigs"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/portalconfigs/default"),
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
