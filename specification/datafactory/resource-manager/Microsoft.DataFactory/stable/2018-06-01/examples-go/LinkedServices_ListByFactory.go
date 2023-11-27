package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/69ece3818b8b0929b43a07c3fe25716427734882/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/LinkedServices_ListByFactory.json
func ExampleLinkedServicesClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLinkedServicesClient().NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.LinkedServiceListResponse = armdatafactory.LinkedServiceListResponse{
		// 	Value: []*armdatafactory.LinkedServiceResource{
		// 		{
		// 			Name: to.Ptr("exampleLinkedService"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/linkedservices"),
		// 			Etag: to.Ptr("0a0064d4-0000-0000-0000-5b245bd00000"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/linkedservices/exampleLinkedService"),
		// 			Properties: &armdatafactory.AzureStorageLinkedService{
		// 				Type: to.Ptr("AzureStorage"),
		// 				Description: to.Ptr("Example description"),
		// 				TypeProperties: &armdatafactory.AzureStorageLinkedServiceTypeProperties{
		// 					ConnectionString: map[string]any{
		// 						"type": "SecureString",
		// 						"value": "**********",
		// 					},
		// 					EncryptedCredential: to.Ptr("ew0KICAiVmVyc2lvbiI6ICIyMDE3LTExLTMwIiwNCiAgIlByb3RlY3Rpb25Nb2RlIjogIktleSIsDQogICJTZWNyZXRDb250ZW50VHlwZSI6ICJQbGFpbnRleHQiLA0KICAiQ3JlZGVudGlhbElkIjogIkRGLURPR0ZPT0QtWUFOWkhBTkctV1VfMGI2M2EyMmYtMGEzNC00NDg2LWIzMDktNzM0NTlkODUyY2Q1Ig0KfQ=="),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
