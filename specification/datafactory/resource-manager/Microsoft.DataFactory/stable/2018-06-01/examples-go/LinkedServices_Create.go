package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/LinkedServices_Create.json
func ExampleLinkedServicesClient_CreateOrUpdate_linkedServicesCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewLinkedServicesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleLinkedService", armdatafactory.LinkedServiceResource{
		Properties: &armdatafactory.AzureStorageLinkedService{
			Type: to.Ptr("AzureStorage"),
			TypeProperties: &armdatafactory.AzureStorageLinkedServiceTypeProperties{
				ConnectionString: map[string]interface{}{
					"type":  "SecureString",
					"value": "DefaultEndpointsProtocol=https;AccountName=examplestorageaccount;AccountKey=<storage key>",
				},
			},
		},
	}, &armdatafactory.LinkedServicesClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
