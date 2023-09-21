package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/ManagedDatabaseMoveOperationResultListMax.json
func ExampleManagedDatabaseMoveOperationsClient_NewListByLocationPager_getsTheLatestManagedDatabaseMoveOperationsForEachDatabaseUnderSpecifiedSubscriptionResourceGroupAndLocationFilteredByOperationType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedDatabaseMoveOperationsClient().NewListByLocationPager("rg1", "westeurope", &armsql.ManagedDatabaseMoveOperationsClientListByLocationOptions{OnlyLatestPerDatabase: nil,
		Filter: to.Ptr("Properties/Operation eq 'StartManagedInstanceDatabaseMove'"),
	})
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
		// page.ManagedDatabaseMoveOperationListResult = armsql.ManagedDatabaseMoveOperationListResult{
		// 	Value: []*armsql.ManagedDatabaseMoveOperationResult{
		// 		{
		// 			Name: to.Ptr("87654321-30a2-f39a-z171-b78695fg32a8"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/managedDatabaseMoveOperationResults"),
		// 			ID: to.Ptr("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1/providers/Microsoft.Sql/locations/westeurope/managedDatabaseMoveOperationResults/87654321-30a2-f39a-z171-b78695fg32a8"),
		// 			Properties: &armsql.ManagedDatabaseMoveOperationResultProperties{
		// 				IsCancellable: to.Ptr(true),
		// 				Operation: to.Ptr("StartManagedInstanceDatabaseMove"),
		// 				OperationFriendlyName: to.Ptr("Start Azure SQL Managed Instance database move"),
		// 				OperationMode: to.Ptr(armsql.MoveOperationModeCopy),
		// 				SourceDatabaseName: to.Ptr("db1"),
		// 				SourceManagedInstanceID: to.Ptr("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1/providers/Microsoft.Sql/managedInstances/mi1"),
		// 				SourceManagedInstanceName: to.Ptr("mi1"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-24T12:54:29.72Z"); return t}()),
		// 				State: to.Ptr(armsql.ManagementOperationStateInProgress),
		// 				TargetDatabaseName: to.Ptr("db1"),
		// 				TargetManagedInstanceID: to.Ptr("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg2/providers/Microsoft.Sql/managedInstances/mi2"),
		// 				TargetManagedInstanceName: to.Ptr("mi2"),
		// 			},
		// 	}},
		// }
	}
}
