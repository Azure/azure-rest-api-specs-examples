package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolColumnSensitivityLabelWithAllParameters.json
func ExampleSQLPoolSensitivityLabelsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSQLPoolSensitivityLabelsClient().CreateOrUpdate(ctx, "myRG", "myServer", "myDatabase", "dbo", "myTable", "myColumn", armsynapse.SensitivityLabel{
		Properties: &armsynapse.SensitivityLabelProperties{
			InformationType:   to.Ptr("PhoneNumber"),
			InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
			LabelID:           to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
			LabelName:         to.Ptr("PII"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SensitivityLabel = armsynapse.SensitivityLabel{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/schemas/tables/columns/sensitivityLabels"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Synapse/workspaces/myServer/sqlPools/myDatabase/schemas/dbo/tables/myTable/columns/myColumn/sensitivityLabels/current"),
	// 	Properties: &armsynapse.SensitivityLabelProperties{
	// 		InformationType: to.Ptr("PhoneNumber"),
	// 		InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
	// 		LabelID: to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
	// 		LabelName: to.Ptr("PII"),
	// 	},
	// }
}
