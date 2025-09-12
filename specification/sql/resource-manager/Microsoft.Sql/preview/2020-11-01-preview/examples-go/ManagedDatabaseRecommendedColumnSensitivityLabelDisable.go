package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8358c7473dfe057d84a6b6a921225063c040b31a/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseRecommendedColumnSensitivityLabelDisable.json
func ExampleManagedDatabaseSensitivityLabelsClient_DisableRecommendation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewManagedDatabaseSensitivityLabelsClient().DisableRecommendation(ctx, "myRG", "myManagedInstanceName", "myDatabase", "dbo", "myTable", "myColumn", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
