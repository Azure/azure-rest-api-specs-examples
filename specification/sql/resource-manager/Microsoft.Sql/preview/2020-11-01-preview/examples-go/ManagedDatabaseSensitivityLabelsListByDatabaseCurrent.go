package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSensitivityLabelsListByDatabaseCurrent.json
func ExampleManagedDatabaseSensitivityLabelsClient_NewListCurrentByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedDatabaseSensitivityLabelsClient().NewListCurrentByDatabasePager("myRG", "myManagedInstanceName", "myDatabase", &armsql.ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseOptions{SkipToken: nil,
		Count:  nil,
		Filter: nil,
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
		// page.SensitivityLabelListResult = armsql.SensitivityLabelListResult{
		// 	Value: []*armsql.SensitivityLabel{
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases/schemas/tables/columns/sensitivityLabels"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/managedInstances/myManagedInstanceName/databases/myDatabase/schemas/dbo/tables/myTable/columns/myColumn/sensitivityLabels/current"),
		// 			Properties: &armsql.SensitivityLabelProperties{
		// 				ColumnName: to.Ptr("myColumn"),
		// 				InformationType: to.Ptr("Financial"),
		// 				InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
		// 				LabelID: to.Ptr("05e6eaa1-075a-4fb4-a732-a92215a2444a"),
		// 				LabelName: to.Ptr("Sensitive"),
		// 				Rank: to.Ptr(armsql.SensitivityLabelRankCritical),
		// 				SchemaName: to.Ptr("dbo"),
		// 				TableName: to.Ptr("myTable"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases/schemas/tables/columns/sensitivityLabels"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/managedInstances/myManagedInstanceName/databases/myDatabase/schemas/dbo/tables/myTable/columns/myColumn3/sensitivityLabels/current"),
		// 			Properties: &armsql.SensitivityLabelProperties{
		// 				ColumnName: to.Ptr("myColumn3"),
		// 				InformationType: to.Ptr("Email"),
		// 				InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
		// 				LabelID: to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
		// 				LabelName: to.Ptr("PII"),
		// 				Rank: to.Ptr(armsql.SensitivityLabelRankHigh),
		// 				SchemaName: to.Ptr("dbo"),
		// 				TableName: to.Ptr("myTable"),
		// 			},
		// 	}},
		// }
	}
}
