package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountSchemas_Get.json
func ExampleIntegrationAccountSchemasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationAccountSchemasClient().Get(ctx, "testResourceGroup", "testIntegrationAccount", "testSchema", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationAccountSchema = armlogic.IntegrationAccountSchema{
	// 	Name: to.Ptr("IntegrationAccountSchema5349"),
	// 	Type: to.Ptr("Microsoft.Logic/integrationAccounts/schemas"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationAccounts/testIntegrationAccount/schemas/testSchema"),
	// 	Properties: &armlogic.IntegrationAccountSchemaProperties{
	// 		ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-24T01:46:48.084Z"); return t}()),
	// 		ContentLink: &armlogic.ContentLink{
	// 			ContentHash: &armlogic.ContentHash{
	// 				Algorithm: to.Ptr("md5"),
	// 				Value: to.Ptr("<value>"),
	// 			},
	// 			ContentSize: to.Ptr[int64](7901),
	// 			ContentVersion: to.Ptr("\"0x8D45C56FEDFCB45\""),
	// 			URI: to.Ptr("<contentLinkUrl>"),
	// 		},
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-24T01:46:48.083Z"); return t}()),
	// 		DocumentName: to.Ptr("OrderFile"),
	// 		Metadata: map[string]any{
	// 		},
	// 		SchemaType: to.Ptr(armlogic.SchemaTypeXML),
	// 		TargetNamespace: to.Ptr("http://Inbound_EDI.OrderFile"),
	// 	},
	// }
}
