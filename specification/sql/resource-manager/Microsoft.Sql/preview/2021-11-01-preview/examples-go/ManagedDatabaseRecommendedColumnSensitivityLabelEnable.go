package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ManagedDatabaseRecommendedColumnSensitivityLabelEnable.json
func ExampleManagedDatabaseSensitivityLabelsClient_EnableRecommendation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsql.NewManagedDatabaseSensitivityLabelsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.EnableRecommendation(ctx,
		"<resource-group-name>",
		"<managed-instance-name>",
		"<database-name>",
		"<schema-name>",
		"<table-name>",
		"<column-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
