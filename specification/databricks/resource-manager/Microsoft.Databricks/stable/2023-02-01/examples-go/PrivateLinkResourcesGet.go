package armdatabricks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e1a87e1a5deb3f986ea1474d233d6680f1e19fc1/specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/PrivateLinkResourcesGet.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabricks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "myResourceGroup", "myWorkspace", "databricks_ui_api", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GroupIDInformation = armdatabricks.GroupIDInformation{
	// 	Name: to.Ptr("databricks_ui_api"),
	// 	Type: to.Ptr("Microsoft.Databricks/workspaces/PrivateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Databricks/workspaces/myWorkspace/PrivateLinkResources/databricks_ui_api"),
	// 	Properties: &armdatabricks.GroupIDInformationProperties{
	// 		GroupID: to.Ptr("databricks_ui_api"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("databricks_ui_api")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.azuredatabricks.net")},
	// 			},
	// 		}
}
