package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateConnection.json
func ExampleConnectionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomation.NewConnectionClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg",
		"myAutomationAccount28",
		"mysConnection",
		armautomation.ConnectionCreateOrUpdateParameters{
			Name: to.Ptr("mysConnection"),
			Properties: &armautomation.ConnectionCreateOrUpdateProperties{
				Description: to.Ptr("my description goes here"),
				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
					Name: to.Ptr("Azure"),
				},
				FieldDefinitionValues: map[string]*string{
					"AutomationCertificateName": to.Ptr("mysCertificateName"),
					"SubscriptionID":            to.Ptr("subid"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
