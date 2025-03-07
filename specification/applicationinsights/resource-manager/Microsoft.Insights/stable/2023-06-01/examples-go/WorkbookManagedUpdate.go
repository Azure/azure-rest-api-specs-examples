package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a0168458930c86636a76bcd7acfdc9c81291bfc/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2023-06-01/examples/WorkbookManagedUpdate.json
func ExampleWorkbooksClient_Update_workbookManagedUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkbooksClient().Update(ctx, "my-resource-group", "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2", &armapplicationinsights.WorkbooksClientUpdateOptions{SourceID: to.Ptr("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group"),
		WorkbookUpdateParameters: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Workbook = armapplicationinsights.Workbook{
	// 	Name: to.Ptr("deadb33f-5e0d-4064-8ebb-1a4ed0313eb2"),
	// 	Type: to.Ptr("Microsoft.Insights/workbooks"),
	// 	ID: to.Ptr("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group/providers/Microsoft.Insights/workbooks/deadb33f-5e0d-4064-8ebb-1a4ed0313eb2"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"TagSample01": to.Ptr("sample01"),
	// 		"TagSample02": to.Ptr("sample02"),
	// 		"hidden-title": to.Ptr("Sample workbook"),
	// 	},
	// 	Identity: &armapplicationinsights.WorkbookResourceIdentity{
	// 		Type: to.Ptr(armapplicationinsights.ManagedServiceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armapplicationinsights.UserAssignedIdentity{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/my-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myid": &armapplicationinsights.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 		},
	// 	},
	// 	Kind: to.Ptr(armapplicationinsights.WorkbookSharedTypeKindShared),
	// 	Properties: &armapplicationinsights.WorkbookProperties{
	// 		Description: to.Ptr("Sample workbook"),
	// 		Category: to.Ptr("workbook"),
	// 		DisplayName: to.Ptr("Sample workbook"),
	// 		Revision: to.Ptr("1e2f8435b98248febee70c64ac22e1bb"),
	// 		SerializedData: to.Ptr("{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":{\"json\":\"test\"},\"name\":\"text - 0\"}],\"isLocked\":false,\"fallbackResourceIds\":[\"/subscriptions/8980832b-9589-4ac2-b322-a6ae6a97f02b/resourceGroups/my-resource-group\"]}"),
	// 		SourceID: to.Ptr("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group"),
	// 		StorageURI: to.Ptr("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourceGroups/my-resource-group/providers/Microsoft.Storage/storageAccounts/mystorage/blobServices/default/containers/mycontainer"),
	// 		TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-20T22:15:08.187Z"); return t}()),
	// 		UserID: to.Ptr("70d90f65-8a70-4e42-b8d5-863725e0a90f"),
	// 		Version: to.Ptr("Notebook/1.0"),
	// 	},
	// }
}
