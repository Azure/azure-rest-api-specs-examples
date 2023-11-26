package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Projects_CreateOrUpdate.json
func ExampleProjectsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectsClient().CreateOrUpdate(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", armdatamigration.Project{
		Location: to.Ptr("southcentralus"),
		Properties: &armdatamigration.ProjectProperties{
			SourcePlatform: to.Ptr(armdatamigration.ProjectSourcePlatformSQL),
			TargetPlatform: to.Ptr(armdatamigration.ProjectTargetPlatformSQLDB),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Project = armdatamigration.Project{
	// 	Name: to.Ptr("DmsSdkProject"),
	// 	Type: to.Ptr("Microsoft.DataMigration/services/projects"),
	// 	ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/DmsSdkProject"),
	// 	Location: to.Ptr("southcentralus"),
	// 	Properties: &armdatamigration.ProjectProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-03T09:14:54.245Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armdatamigration.ProjectProvisioningStateSucceeded),
	// 		SourcePlatform: to.Ptr(armdatamigration.ProjectSourcePlatformSQL),
	// 		TargetPlatform: to.Ptr(armdatamigration.ProjectTargetPlatformSQLDB),
	// 	},
	// }
}
