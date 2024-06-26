package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getConnectionType.json
func ExampleConnectionTypeClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionTypeClient().Get(ctx, "rg", "myAutomationAccount22", "myCT", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectionType = armautomation.ConnectionType{
	// 	Name: to.Ptr("myCT"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount22/connectionTypes/myCT"),
	// 	Properties: &armautomation.ConnectionTypeProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T03:47:53.433Z"); return t}()),
	// 		FieldDefinitions: map[string]*armautomation.FieldDefinition{
	// 			"myBoolField": &armautomation.FieldDefinition{
	// 				Type: to.Ptr("bool"),
	// 				IsEncrypted: to.Ptr(false),
	// 				IsOptional: to.Ptr(false),
	// 			},
	// 			"myStringField": &armautomation.FieldDefinition{
	// 				Type: to.Ptr("string"),
	// 				IsEncrypted: to.Ptr(false),
	// 				IsOptional: to.Ptr(false),
	// 			},
	// 			"myStringFieldEncrypted": &armautomation.FieldDefinition{
	// 				Type: to.Ptr("string"),
	// 				IsEncrypted: to.Ptr(true),
	// 				IsOptional: to.Ptr(false),
	// 			},
	// 		},
	// 		IsGlobal: to.Ptr(false),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T03:47:53.557Z"); return t}()),
	// 	},
	// }
}
