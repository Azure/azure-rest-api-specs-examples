package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-01-01-preview/examples/Bindings_CreateOrUpdate.json
func ExampleBindingsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBindingsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myservice", "myapp", "mybinding", armappplatform.BindingResource{
		Properties: &armappplatform.BindingResourceProperties{
			BindingParameters: map[string]*string{
				"apiType":      to.Ptr("SQL"),
				"databaseName": to.Ptr("db1"),
			},
			Key:        to.Ptr("xxxx"),
			ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DocumentDB/databaseAccounts/my-cosmosdb-1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BindingResource = armappplatform.BindingResource{
	// 	Name: to.Ptr("mybinding"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/apps/bindings"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/apps/myapp/bindings/mybinding"),
	// 	SystemData: &armappplatform.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample-user"),
	// 		CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample-user"),
	// 		LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armappplatform.BindingResourceProperties{
	// 		BindingParameters: map[string]*string{
	// 			"apiType": to.Ptr("SQL"),
	// 			"databaseName": to.Ptr("db1"),
	// 		},
	// 		CreatedAt: to.Ptr("2019-01-01T12:34:56.000Z"),
	// 		GeneratedProperties: to.Ptr("<connection-string>"),
	// 		ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DocumentDB/databaseAccounts/my-cosmosdb-1"),
	// 		ResourceName: to.Ptr("my-cosmosdb-1"),
	// 		ResourceType: to.Ptr("Microsoft.DocumentDB"),
	// 		UpdatedAt: to.Ptr("2019-01-01T12:34:56.000Z"),
	// 	},
	// }
}
