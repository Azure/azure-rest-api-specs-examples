package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ColumnSensitivityLabelCreateMax.json
func ExampleSensitivityLabelsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSensitivityLabelsClient().CreateOrUpdate(ctx, "myRG", "myServer", "myDatabase", "dbo", "myTable", "myColumn", armsql.SensitivityLabel{
		Properties: &armsql.SensitivityLabelProperties{
			InformationType:   to.Ptr("PhoneNumber"),
			InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
			LabelID:           to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
			LabelName:         to.Ptr("PII"),
			Rank:              to.Ptr(armsql.SensitivityLabelRankLow),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SensitivityLabel = armsql.SensitivityLabel{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns/sensitivityLabels"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/myServer/databases/myDatabase/schemas/dbo/tables/myTable/columns/myColumn/sensitivityLabels/current"),
	// 	Properties: &armsql.SensitivityLabelProperties{
	// 		ColumnName: to.Ptr("myColumn"),
	// 		InformationType: to.Ptr("PhoneNumber"),
	// 		InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
	// 		LabelID: to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
	// 		LabelName: to.Ptr("PII"),
	// 		Rank: to.Ptr(armsql.SensitivityLabelRankMedium),
	// 		SchemaName: to.Ptr("dbo"),
	// 		TableName: to.Ptr("myTable"),
	// 	},
	// }
}
