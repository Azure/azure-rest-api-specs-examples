package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/LinkedServices_Get.json
func ExampleLinkedServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLinkedServicesClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleLinkedService", &armdatafactory.LinkedServicesClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LinkedServiceResource = armdatafactory.LinkedServiceResource{
	// 	Name: to.Ptr("exampleLinkedService"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/linkedservices"),
	// 	Etag: to.Ptr("1500474f-0000-0200-0000-5cbe090d0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/linkedservices/exampleLinkedService"),
	// 	Properties: &armdatafactory.AzureStorageLinkedService{
	// 		Type: to.Ptr("AzureStorage"),
	// 		Description: to.Ptr("Example description"),
	// 		TypeProperties: &armdatafactory.AzureStorageLinkedServiceTypeProperties{
	// 			ConnectionString: map[string]any{
	// 				"type": "SecureString",
	// 				"value": "**********",
	// 			},
	// 			EncryptedCredential: to.Ptr("ew0KICAiVmVyc2lvbiI6ICIyMDE3LTExLTMwIiwNCiAgIlByb3RlY3Rpb25Nb2RlIjogIktleSIsDQogICJTZWNyZXRDb250ZW50VHlwZSI6ICJQbGFpbnRleHQiLA0KICAiQ3JlZGVudGlhbElkIjogIkhWVEVTVEVYQU1QTEVTXzg5MjU1NGY0LTViNGItNDFhOS1hYWYxLTg4ZWI5ZDBjOWIzMyINCn0="),
	// 		},
	// 	},
	// }
}
