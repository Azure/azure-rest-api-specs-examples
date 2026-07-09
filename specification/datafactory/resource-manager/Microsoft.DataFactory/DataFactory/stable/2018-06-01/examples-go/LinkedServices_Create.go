package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v11"
)

// Generated from example definition: 2018-06-01/LinkedServices_Create.json
func ExampleLinkedServicesClient_CreateOrUpdate_linkedServicesCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLinkedServicesClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleLinkedService", armdatafactory.LinkedServiceResource{
		Properties: &armdatafactory.AzureStorageLinkedService{
			Type: to.Ptr("AzureStorage"),
			TypeProperties: &armdatafactory.AzureStorageLinkedServiceTypeProperties{
				ConnectionString: map[string]any{
					"type":  "SecureString",
					"value": "DefaultEndpointsProtocol=https;AccountName=examplestorageaccount;AccountKey=<storage key>",
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatafactory.LinkedServicesClientCreateOrUpdateResponse{
	// 	LinkedServiceResource: armdatafactory.LinkedServiceResource{
	// 		Name: to.Ptr("exampleLinkedService"),
	// 		Type: to.Ptr("Microsoft.DataFactory/factories/linkedservices"),
	// 		Etag: to.Ptr("0a0062d4-0000-0000-0000-5b245bcf0000"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/linkedservices/exampleLinkedService"),
	// 		Properties: &armdatafactory.AzureStorageLinkedService{
	// 			Type: to.Ptr("AzureStorage"),
	// 			TypeProperties: &armdatafactory.AzureStorageLinkedServiceTypeProperties{
	// 				ConnectionString: map[string]any{
	// 					"type": "SecureString",
	// 					"value": "**********",
	// 				},
	// 				EncryptedCredential: to.Ptr("ew0KICAiVmVyc2lvbiI6ICIyMDE3LTExLTMwIiwNCiAgIlByb3RlY3Rpb25Nb2RlIjogIktleSIsDQogICJTZWNyZXRDb250ZW50VHlwZSI6ICJQbGFpbnRleHQiLA0KICAiQ3JlZGVudGlhbElkIjogIkRGLURPR0ZPT0QtWUFOWkhBTkctV1VfM2FiMTk0NjYtNWUxNi00NzU1LWJlNzktMjI2ZTVmZWU3YzY0Ig0KfQ=="),
	// 			},
	// 		},
	// 	},
	// }
}
