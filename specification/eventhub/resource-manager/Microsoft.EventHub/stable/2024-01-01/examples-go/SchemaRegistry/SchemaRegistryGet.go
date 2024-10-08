package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b574e2a41acda14a90ef237006e8bbdda2b63c63/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/SchemaRegistry/SchemaRegistryGet.json
func ExampleSchemaRegistryClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSchemaRegistryClient().Get(ctx, "alitest", "ali-ua-test-eh-system-1", "testSchemaGroup1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SchemaGroup = armeventhub.SchemaGroup{
	// 	Name: to.Ptr("testSchemaGroup1"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces/SchemaGroups"),
	// 	ID: to.Ptr("/subscriptions/e8baea74-64ce-459b-bee3-5aa4c47b3ae3/resourceGroups/alitest/providers/Microsoft.EventHub/namespaces/ali-ua-test-eh-system-1/schemagroups/testSchemaGroup1"),
	// 	Location: to.Ptr("EAST US 2 EUAP"),
	// 	Properties: &armeventhub.SchemaGroupProperties{
	// 		CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-13T03:08:11.167Z"); return t}()),
	// 		ETag: to.Ptr("51ddcff4-a287-423c-b194-7a8ebbfd8366"),
	// 		GroupProperties: map[string]*string{
	// 		},
	// 		SchemaCompatibility: to.Ptr(armeventhub.SchemaCompatibilityForward),
	// 		SchemaType: to.Ptr(armeventhub.SchemaTypeAvro),
	// 		UpdatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-13T03:08:11.167Z"); return t}()),
	// 	},
	// }
}
