package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/Projects_List.json
func ExampleProjectsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProjectsClient().NewListPager("DmsSdkRg", "DmsSdkService", nil)
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
		// page.ProjectList = armdatamigration.ProjectList{
		// 	Value: []*armdatamigration.Project{
		// 		{
		// 			Name: to.Ptr("project1"),
		// 			Type: to.Ptr("Microsoft.DataMigration/services/projects"),
		// 			ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/project1"),
		// 			Location: to.Ptr("southcentralus"),
		// 			Properties: &armdatamigration.ProjectProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-14T01:29:56.304Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armdatamigration.ProjectProvisioningStateSucceeded),
		// 				SourcePlatform: to.Ptr(armdatamigration.ProjectSourcePlatformSQL),
		// 				TargetPlatform: to.Ptr(armdatamigration.ProjectTargetPlatformSQLDB),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("project2"),
		// 			Type: to.Ptr("Microsoft.DataMigration/services/projects"),
		// 			ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/project2"),
		// 			Location: to.Ptr("southcentralus"),
		// 			Properties: &armdatamigration.ProjectProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-14T01:30:05.618Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armdatamigration.ProjectProvisioningStateSucceeded),
		// 				SourcePlatform: to.Ptr(armdatamigration.ProjectSourcePlatformSQL),
		// 				TargetPlatform: to.Ptr(armdatamigration.ProjectTargetPlatformSQLDB),
		// 			},
		// 	}},
		// }
	}
}
