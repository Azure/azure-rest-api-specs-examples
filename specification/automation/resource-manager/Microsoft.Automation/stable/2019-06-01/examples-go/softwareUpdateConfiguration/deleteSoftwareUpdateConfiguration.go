package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/deleteSoftwareUpdateConfiguration.json
func ExampleSoftwareUpdateConfigurationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomation.NewSoftwareUpdateConfigurationsClient("51766542-3ed7-4a72-a187-0c8ab644ddab", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"mygroup",
		"myaccount",
		"mypatch",
		&armautomation.SoftwareUpdateConfigurationsClientDeleteOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
