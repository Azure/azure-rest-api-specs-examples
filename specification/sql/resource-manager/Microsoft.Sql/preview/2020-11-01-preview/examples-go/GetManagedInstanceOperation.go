package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetManagedInstanceOperation.json
func ExampleManagedInstanceOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedInstanceOperationsClient().Get(ctx, "sqlcrudtest-7398", "sqlcrudtest-4645", "00000000-1111-2222-3333-444444444444", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedInstanceOperation = armsql.ManagedInstanceOperation{
	// 	Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/operations"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/managedInstances/sqlcrudtest-4645/operations/11111111-1111-1111-1111-111111111111"),
	// 	Properties: &armsql.ManagedInstanceOperationProperties{
	// 		IsCancellable: to.Ptr(false),
	// 		ManagedInstanceName: to.Ptr("sqlcrudtest-4645"),
	// 		Operation: to.Ptr("UpsertManagedServer"),
	// 		OperationFriendlyName: to.Ptr("UPDATE MANAGED SERVER"),
	// 		OperationSteps: &armsql.ManagedInstanceOperationSteps{
	// 			CurrentStep: to.Ptr[int32](2),
	// 			StepsList: []*armsql.UpsertManagedServerOperationStep{
	// 				{
	// 					Name: to.Ptr("Request validation"),
	// 					Order: to.Ptr[int32](1),
	// 					Status: to.Ptr(armsql.UpsertManagedServerOperationStepStatusCompleted),
	// 				},
	// 				{
	// 					Name: to.Ptr("New SQL Instance Startup"),
	// 					Order: to.Ptr[int32](2),
	// 					Status: to.Ptr(armsql.UpsertManagedServerOperationStepStatusCanceled),
	// 				},
	// 				{
	// 					Name: to.Ptr("Seeding database files"),
	// 					Order: to.Ptr[int32](3),
	// 					Status: to.Ptr(armsql.UpsertManagedServerOperationStepStatusNotStarted),
	// 				},
	// 				{
	// 					Name: to.Ptr("Preparing Failover and Failover"),
	// 					Order: to.Ptr[int32](4),
	// 					Status: to.Ptr(armsql.UpsertManagedServerOperationStepStatusNotStarted),
	// 				},
	// 				{
	// 					Name: to.Ptr("Old SQL Instance cleanup"),
	// 					Order: to.Ptr[int32](5),
	// 					Status: to.Ptr(armsql.UpsertManagedServerOperationStepStatusNotStarted),
	// 			}},
	// 			TotalSteps: to.Ptr("5"),
	// 		},
	// 		PercentComplete: to.Ptr[int32](100),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-06T11:08:44.49Z"); return t}()),
	// 		State: to.Ptr(armsql.ManagementOperationStateCancelled),
	// 	},
	// }
}
