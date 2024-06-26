package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SensitivityLabelsRecommendedUpdate.json
func ExampleSQLPoolRecommendedSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSQLPoolRecommendedSensitivityLabelsClient().Update(ctx, "myRG", "myWorkspace", "mySqlPool", armsynapse.RecommendedSensitivityLabelUpdateList{
		Operations: []*armsynapse.RecommendedSensitivityLabelUpdate{
			{
				Properties: &armsynapse.RecommendedSensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("column1"),
					Op:     to.Ptr(armsynapse.RecommendedSensitivityLabelUpdateKindEnable),
					Table:  to.Ptr("table1"),
				},
			},
			{
				Properties: &armsynapse.RecommendedSensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("column2"),
					Op:     to.Ptr(armsynapse.RecommendedSensitivityLabelUpdateKindEnable),
					Table:  to.Ptr("table2"),
				},
			},
			{
				Properties: &armsynapse.RecommendedSensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("column3"),
					Op:     to.Ptr(armsynapse.RecommendedSensitivityLabelUpdateKindDisable),
					Table:  to.Ptr("table1"),
				},
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
