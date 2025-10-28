package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry/v2"
)

// Generated from example definition: 2025-10-01/List_Schemas_BySchemaRegistry.json
func ExampleSchemasClient_NewListBySchemaRegistryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSchemasClient().NewListBySchemaRegistryPager("myResourceGroup", "my-schema-registry", nil)
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
		// page = armdeviceregistry.SchemasClientListBySchemaRegistryResponse{
		// 	SchemaListResult: armdeviceregistry.SchemaListResult{
		// 		Value: []*armdeviceregistry.Schema{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/schemaRegistries/my-schema-registry/schemas/my-schema"),
		// 				Name: to.Ptr("my-schema"),
		// 				Type: to.Ptr("Microsoft.DeviceRegistry/schemaRegistries/schemas"),
		// 				SystemData: &armdeviceregistry.SystemData{
		// 					CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
		// 				},
		// 				Properties: &armdeviceregistry.SchemaProperties{
		// 					UUID: to.Ptr("0796f7c1-f2c8-44d7-9f5b-9a6f9522a85d"),
		// 					DisplayName: to.Ptr("My Schema"),
		// 					Description: to.Ptr("This is a sample Schema"),
		// 					Format: to.Ptr(armdeviceregistry.FormatJSONSchemaDraft7),
		// 					SchemaType: to.Ptr(armdeviceregistry.SchemaTypeMessageSchema),
		// 					ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 					Tags: map[string]*string{
		// 						"sampleKey": to.Ptr("sampleValue"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/schemaRegistries/my-schema-registry/schemas/my-schema-2"),
		// 				Name: to.Ptr("my-schema-2"),
		// 				Type: to.Ptr("Microsoft.DeviceRegistry/schemaRegistries/schemas"),
		// 				SystemData: &armdeviceregistry.SystemData{
		// 					CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
		// 				},
		// 				Properties: &armdeviceregistry.SchemaProperties{
		// 					UUID: to.Ptr("7824a74f-21e1-4458-ae06-604d3a241d2c"),
		// 					DisplayName: to.Ptr("My Schema 2"),
		// 					Description: to.Ptr("This is another sample Schema"),
		// 					Format: to.Ptr(armdeviceregistry.FormatJSONSchemaDraft7),
		// 					SchemaType: to.Ptr(armdeviceregistry.SchemaTypeMessageSchema),
		// 					ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 					Tags: map[string]*string{
		// 						"sampleKey": to.Ptr("sampleValue"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
