package armappplatform_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-03-01-preview/examples/ConfigurationServices_Validate.json
func ExampleConfigurationServicesClient_BeginValidate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armappplatform.NewConfigurationServicesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginValidate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<configuration-service-name>",
		armappplatform.ConfigurationServiceSettings{
			GitProperty: &armappplatform.ConfigurationServiceGitProperty{
				Repositories: []*armappplatform.ConfigurationServiceGitRepository{
					{
						Name:  to.Ptr("<name>"),
						Label: to.Ptr("<label>"),
						Patterns: []*string{
							to.Ptr("app/dev")},
						URI: to.Ptr("<uri>"),
					}},
			},
		},
		&armappplatform.ConfigurationServicesClientBeginValidateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
