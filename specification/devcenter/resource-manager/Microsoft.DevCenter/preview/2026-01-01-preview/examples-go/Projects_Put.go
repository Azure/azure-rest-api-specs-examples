package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/Projects_Put.json
func ExampleProjectsClient_BeginCreateOrUpdate_projectsCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProjectsClient().BeginCreateOrUpdate(ctx, "rg1", "DevProject", armdevcenter.Project{
		Location: to.Ptr("centralus"),
		Properties: &armdevcenter.ProjectProperties{
			Description: to.Ptr("This is my first project."),
			DevCenterID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso"),
			DisplayName: to.Ptr("Dev"),
			AssignedGroups: []*armdevcenter.AssignedGroup{
				{
					ObjectID: to.Ptr("11111111-1111-1111-1111-111111111111"),
					Scope:    to.Ptr(armdevcenter.AssignedGroupScopeDevBox),
				},
				{
					ObjectID: to.Ptr("22222222-2222-2222-2222-222222222222"),
					Scope:    to.Ptr(armdevcenter.AssignedGroupScopeDevBox),
				},
			},
		},
		Tags: map[string]*string{
			"CostCenter": to.Ptr("R&D"),
		},
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
	// res = armdevcenter.ProjectsClientCreateOrUpdateResponse{
	// 	Project: armdevcenter.Project{
	// 		Name: to.Ptr("DevProject"),
	// 		Type: to.Ptr("Microsoft.DevCenter/projects"),
	// 		ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject"),
	// 		Location: to.Ptr("centralus"),
	// 		Properties: &armdevcenter.ProjectProperties{
	// 			Description: to.Ptr("This is my first project."),
	// 			DevCenterID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso"),
	// 			DevCenterURI: to.Ptr("https://4c7c8922-78e9-4928-aa6f-75ba59355371-contoso.centralus.devcenter.azure.com"),
	// 			AssignedGroups: []*armdevcenter.AssignedGroup{
	// 				{
	// 					ObjectID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					Scope: to.Ptr(armdevcenter.AssignedGroupScopeDevBox),
	// 				},
	// 				{
	// 					ObjectID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 					Scope: to.Ptr(armdevcenter.AssignedGroupScopeDevBox),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armdevcenter.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user1"),
	// 			LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"CostCenter": to.Ptr("R&D"),
	// 			"hidden-title": to.Ptr("Dev"),
	// 		},
	// 	},
	// }
}
