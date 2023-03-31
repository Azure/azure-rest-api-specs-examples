package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolWorkloadGroupWorkloadClassifierList.json
func ExampleSQLPoolWorkloadClassifierClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLPoolWorkloadClassifierClient().NewListPager("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187", "wlm_workloadgroup", nil)
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
		// page.WorkloadClassifierListResult = armsynapse.WorkloadClassifierListResult{
		// 	Value: []*armsynapse.WorkloadClassifier{
		// 		{
		// 			Name: to.Ptr("classifier3"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/workloadGroups/workloadClassifiers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Synapse/workspaces/sqlcrudtest-2080/sqlPools/sqlcrudtest-9187/workloadGroups/wlm_workloadgroup/workloadClassifiers/classifier3"),
		// 			Properties: &armsynapse.WorkloadClassifierProperties{
		// 				Context: to.Ptr(""),
		// 				EndTime: to.Ptr(""),
		// 				Importance: to.Ptr("high"),
		// 				Label: to.Ptr(""),
		// 				MemberName: to.Ptr("dbo"),
		// 				StartTime: to.Ptr(""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("classifier1"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/workloadGroups/workloadClassifiers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Synapse/workspaces/sqlcrudtest-2080/sqlPools/sqlcrudtest-9187/workloadGroups/wlm_workloadgroup/workloadClassifiers/classifier1"),
		// 			Properties: &armsynapse.WorkloadClassifierProperties{
		// 				Context: to.Ptr("test_context"),
		// 				EndTime: to.Ptr("14:00"),
		// 				Importance: to.Ptr("high"),
		// 				Label: to.Ptr("test_label"),
		// 				MemberName: to.Ptr("dbo"),
		// 				StartTime: to.Ptr("12:00"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("classifier2"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/workloadGroups/workloadClassifiers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Synapse/workspaces/sqlcrudtest-2080/sqlPools/sqlcrudtest-9187/workloadGroups/wlm_workloadgroup/workloadClassifiers/classifier2"),
		// 			Properties: &armsynapse.WorkloadClassifierProperties{
		// 				Context: to.Ptr(""),
		// 				EndTime: to.Ptr("17:00"),
		// 				Importance: to.Ptr("high"),
		// 				Label: to.Ptr(""),
		// 				MemberName: to.Ptr("dbo"),
		// 				StartTime: to.Ptr("11:00"),
		// 			},
		// 	}},
		// }
	}
}
