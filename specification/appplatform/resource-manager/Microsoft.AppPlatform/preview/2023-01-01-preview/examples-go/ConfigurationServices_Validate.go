package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-01-01-preview/examples/ConfigurationServices_Validate.json
func ExampleConfigurationServicesClient_BeginValidate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConfigurationServicesClient().BeginValidate(ctx, "myResourceGroup", "myservice", "default", armappplatform.ConfigurationServiceSettings{
		GitProperty: &armappplatform.ConfigurationServiceGitProperty{
			Repositories: []*armappplatform.ConfigurationServiceGitRepository{
				{
					Name:  to.Ptr("fake"),
					Label: to.Ptr("master"),
					Patterns: []*string{
						to.Ptr("app/dev")},
					URI: to.Ptr("https://github.com/fake-user/fake-repository"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationServiceSettingsValidateResult = armappplatform.ConfigurationServiceSettingsValidateResult{
	// 	GitPropertyValidationResult: &armappplatform.ConfigurationServiceGitPropertyValidateResult{
	// 		IsValid: to.Ptr(true),
	// 	},
	// }
}
