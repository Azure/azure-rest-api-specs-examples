package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateWorkloadClassifierMax.json
func ExampleWorkloadClassifiersClient_BeginCreateOrUpdate_createAWorkloadGroupWithAllPropertiesSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadClassifiersClient().BeginCreateOrUpdate(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", "wlm_workloadgroup", "wlm_workloadclassifier", armsql.WorkloadClassifier{
		Properties: &armsql.WorkloadClassifierProperties{
			Context:    to.Ptr("test_context"),
			EndTime:    to.Ptr("14:00"),
			Importance: to.Ptr("high"),
			Label:      to.Ptr("test_label"),
			MemberName: to.Ptr("dbo"),
			StartTime:  to.Ptr("12:00"),
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
	// res.WorkloadClassifier = armsql.WorkloadClassifier{
	// 	Name: to.Ptr("wlm_workloadclassifier"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/workloadGroups/workloadClassifiers"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/workloadGroups/wlm_workloadgroup/workloadClassifiers/wlm_workloadclassifier"),
	// 	Properties: &armsql.WorkloadClassifierProperties{
	// 		Context: to.Ptr("test_context"),
	// 		EndTime: to.Ptr("14:00"),
	// 		Importance: to.Ptr("high"),
	// 		Label: to.Ptr("test_label"),
	// 		MemberName: to.Ptr("dbo"),
	// 		StartTime: to.Ptr("12:00"),
	// 	},
	// }
}
