package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub/v2"
)

// Generated from example definition: 2026-07-01-preview/SchemaRegistry/SchemaRegistryListByNamespace.json
func ExampleSchemaRegistryClient_NewListByNamespacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("e8baea74-64ce-459b-bee3-5aa4c47b3ae3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSchemaRegistryClient().NewListByNamespacePager("alitest", "ali-ua-test-eh-system-1", nil)
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
		// page = armeventhub.SchemaRegistryClientListByNamespaceResponse{
		// 	SchemaGroupListResult: armeventhub.SchemaGroupListResult{
		// 		Value: []*armeventhub.SchemaGroup{
		// 			{
		// 				Name: to.Ptr("testSchemaGroup1"),
		// 				Type: to.Ptr("Microsoft.EventHub/Namespaces/SchemaGroups"),
		// 				ID: to.Ptr("/subscriptions/e8baea74-64ce-459b-bee3-5aa4c47b3ae3/resourceGroups/alitest/providers/Microsoft.EventHub/namespaces/ali-ua-test-eh-system-1/schemagroups/testSchemaGroup1"),
		// 				Location: to.Ptr("EAST US 2 EUAP"),
		// 				Properties: &armeventhub.SchemaGroupProperties{
		// 					CreatedAtUTC: to.Ptr(time.Date(2021, time.October, 13, 3, 8, 11, 167187900, time.UTC)),
		// 					ETag: to.Ptr("51ddcff4-a287-423c-b194-7a8ebbfd8366"),
		// 					GroupProperties: map[string]*string{
		// 					},
		// 					SchemaCompatibility: to.Ptr(armeventhub.SchemaCompatibilityForward),
		// 					SchemaType: to.Ptr(armeventhub.SchemaTypeAvro),
		// 					UpdatedAtUTC: to.Ptr(time.Date(2021, time.October, 13, 3, 8, 11, 167187900, time.UTC)),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testSchemaGroup2"),
		// 				Type: to.Ptr("Microsoft.EventHub/Namespaces/SchemaGroups"),
		// 				ID: to.Ptr("/subscriptions/e8baea74-64ce-459b-bee3-5aa4c47b3ae3/resourceGroups/alitest/providers/Microsoft.EventHub/namespaces/ali-ua-test-eh-system-1/schemagroups/testSchemaGroup2"),
		// 				Location: to.Ptr("EAST US 2 EUAP"),
		// 				Properties: &armeventhub.SchemaGroupProperties{
		// 					CreatedAtUTC: to.Ptr(time.Date(2021, time.October, 13, 3, 10, 33, 697431900, time.UTC)),
		// 					ETag: to.Ptr("d01173a4-08c5-43c9-b30f-d9666196a907"),
		// 					GroupProperties: map[string]*string{
		// 					},
		// 					SchemaCompatibility: to.Ptr(armeventhub.SchemaCompatibilityNone),
		// 					SchemaType: to.Ptr(armeventhub.SchemaTypeAvro),
		// 					UpdatedAtUTC: to.Ptr(time.Date(2021, time.October, 13, 3, 10, 33, 697431900, time.UTC)),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testSchemaGroup3"),
		// 				Type: to.Ptr("Microsoft.EventHub/Namespaces/SchemaGroups"),
		// 				ID: to.Ptr("/subscriptions/e8baea74-64ce-459b-bee3-5aa4c47b3ae3/resourceGroups/alitest/providers/Microsoft.EventHub/namespaces/ali-ua-test-eh-system-1/schemagroups/testSchemaGroup3"),
		// 				Location: to.Ptr("EAST US 2 EUAP"),
		// 				Properties: &armeventhub.SchemaGroupProperties{
		// 					CreatedAtUTC: to.Ptr(time.Date(2021, time.October, 13, 3, 13, 30, 893858500, time.UTC)),
		// 					ETag: to.Ptr("2c1c3d08-2fb8-4a4e-91f4-6e8b940c1b7c"),
		// 					GroupProperties: map[string]*string{
		// 					},
		// 					SchemaCompatibility: to.Ptr(armeventhub.SchemaCompatibilityBackward),
		// 					SchemaType: to.Ptr(armeventhub.SchemaTypeAvro),
		// 					UpdatedAtUTC: to.Ptr(time.Date(2021, time.October, 13, 3, 13, 30, 893858500, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
