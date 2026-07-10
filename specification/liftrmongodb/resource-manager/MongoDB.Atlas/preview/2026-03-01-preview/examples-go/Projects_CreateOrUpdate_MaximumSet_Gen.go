package armmongodbatlas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongodbatlas/armmongodbatlas"
)

// Generated from example definition: 2026-03-01-preview/Projects_CreateOrUpdate_MaximumSet_Gen.json
func ExampleProjectsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongodbatlas.NewClientFactory("1E4BD993-6890-4E69-8966-81482D7502EF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProjectsClient().BeginCreateOrUpdate(ctx, "rgopenapi", "myOrganization", "myProject", armmongodbatlas.Project{
		Properties: &armmongodbatlas.ProjectProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmongodbatlas.ProjectsClientCreateOrUpdateResponse{
	// 	Project: armmongodbatlas.Project{
	// 		ID: to.Ptr("/subscriptions/1E4BD993-6890-4E69-8966-81482D7502EF/resourceGroups/rgopenapi/providers/MongoDB.Atlas/organizations/myOrganization/projects/myProject"),
	// 		Name: to.Ptr("myProject"),
	// 		Type: to.Ptr("MongoDB.Atlas/organizations/projects"),
	// 		Properties: &armmongodbatlas.ProjectProperties{
	// 			ProjectID: to.Ptr("507f1f77bcf86cd799439011"),
	// 			ProjectName: to.Ptr("myProject"),
	// 			OrganizationID: to.Ptr("507f1f77bcf86cd799439011"),
	// 			ClusterCount: to.Ptr[int64](0),
	// 			ProvisioningState: to.Ptr(armmongodbatlas.ResourceProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armmongodbatlas.SystemData{
	// 			CreatedBy: to.Ptr("rxmwavptbhnfjbl"),
	// 			CreatedByType: to.Ptr(armmongodbatlas.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-18T10:44:11.098Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("vvyjmiqdxlzkubu"),
	// 			LastModifiedByType: to.Ptr(armmongodbatlas.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-18T10:44:11.099Z"); return t}()),
	// 		},
	// 	},
	// }
}
