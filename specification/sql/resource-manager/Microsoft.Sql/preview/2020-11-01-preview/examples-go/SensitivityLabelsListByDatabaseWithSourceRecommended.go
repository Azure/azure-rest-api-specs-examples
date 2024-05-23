package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SensitivityLabelsListByDatabaseWithSourceRecommended.json
func ExampleSensitivityLabelsClient_NewListRecommendedByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSensitivityLabelsClient().NewListRecommendedByDatabasePager("myRG", "myServer", "myDatabase", &armsql.SensitivityLabelsClientListRecommendedByDatabaseOptions{SkipToken: nil,
		IncludeDisabledRecommendations: nil,
		Filter:                         nil,
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
		// 			Name: to.Ptr("recommended"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns/sensitivityLabels"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/myServer/databases/myDatabase/schemas/dbo/tables/myTable/columns/myColumn/sensitivityLabels/recommended"),
		// 			Properties: &armsql.SensitivityLabelProperties{
		// 				ColumnName: to.Ptr("myColumn"),
		// 				InformationType: to.Ptr("Financial"),
		// 				InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
		// 				LabelID: to.Ptr("05e6eaa1-075a-4fb4-a732-a92215a2444a"),
		// 				LabelName: to.Ptr("Sensitive"),
		// 				SchemaName: to.Ptr("dbo"),
		// 				TableName: to.Ptr("myTable"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("recommended"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns/sensitivityLabels"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/myServer/databases/myDatabase/schemas/dbo/tables/myTable/columns/myColumn2/sensitivityLabels/recommended"),
		// 			Properties: &armsql.SensitivityLabelProperties{
		// 				ColumnName: to.Ptr("myColumn2"),
		// 				InformationType: to.Ptr("Email"),
		// 				InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
		// 				LabelID: to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
		// 				LabelName: to.Ptr("PII"),
		// 				SchemaName: to.Ptr("dbo"),
		// 				TableName: to.Ptr("myTable"),
		// 			},
		// 	}},
		// }
	}
}
