package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v3"
)

// Generated from example definition: 2026-06-01/GoldenGateConnections_ListAssignedDeploymentsByParent_MaximumSet_Gen.json
func ExampleGoldenGateConnectionsClient_NewListAssignedDeploymentsByParentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGoldenGateConnectionsClient().NewListAssignedDeploymentsByParentPager("rgopenapi", "resource1", nil)
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
		// page = armoracledatabase.GoldenGateConnectionsClientListAssignedDeploymentsByParentResponse{
		// 	AssignedDeploymentListResult: armoracledatabase.AssignedDeploymentListResult{
		// 		Value: []*armoracledatabase.AssignedDeployment{
		// 			{
		// 				Properties: &armoracledatabase.DeploymentConnectionAssignmentProperties{
		// 					ConnectionID: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					DeploymentID: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					CompartmentID: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					DeploymentName: to.Ptr("resource1"),
		// 					ConnectionName: to.Ptr("resource1"),
		// 					Ocid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					LifecycleState: to.Ptr(armoracledatabase.GoldenGateConnectionAssignmentLifecycleStateCreating),
		// 					TimeCreated: to.Ptr("vwplju"),
		// 					TimeUpdated: to.Ptr("2026-06-01T00:00:00Z"),
		// 					AliasName: to.Ptr("resource1"),
		// 					ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/resources/resource1"),
		// 				Name: to.Ptr("resource1"),
		// 				Type: to.Ptr("Oracle.Database/resource"),
		// 				SystemData: &armoracledatabase.SystemData{
		// 					CreatedBy: to.Ptr("ns"),
		// 					CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2026, time.July, 28, 21, 47, 8, 141000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("example"),
		// 					LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2026, time.July, 28, 21, 47, 8, 141000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
