package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateConnectionType.json
func ExampleConnectionTypeClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewConnectionTypeClient().CreateOrUpdate(ctx, "rg", "myAutomationAccount22", "myCT", armautomation.ConnectionTypeCreateOrUpdateParameters{
		Name: to.Ptr("myCT"),
		Properties: &armautomation.ConnectionTypeCreateOrUpdateProperties{
			FieldDefinitions: map[string]*armautomation.FieldDefinition{
				"myBoolField": {
					Type:        to.Ptr("bool"),
					IsEncrypted: to.Ptr(false),
					IsOptional:  to.Ptr(false),
				},
				"myStringField": {
					Type:        to.Ptr("string"),
					IsEncrypted: to.Ptr(false),
					IsOptional:  to.Ptr(false),
				},
				"myStringFieldEncrypted": {
					Type:        to.Ptr("string"),
					IsEncrypted: to.Ptr(true),
					IsOptional:  to.Ptr(false),
				},
			},
			IsGlobal: to.Ptr(false),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
