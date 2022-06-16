package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementPortalSettingsUpdateSignUp.json
func ExampleSignUpSettingsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewSignUpSettingsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Update(ctx,
		"rg1",
		"apimService1",
		"*",
		armapimanagement.PortalSignupSettings{
			Properties: &armapimanagement.PortalSignupSettingsProperties{
				Enabled: to.Ptr(true),
				TermsOfService: &armapimanagement.TermsOfServiceProperties{
					ConsentRequired: to.Ptr(true),
					Enabled:         to.Ptr(true),
					Text:            to.Ptr("Terms of service text."),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
