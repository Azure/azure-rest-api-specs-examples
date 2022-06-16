package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/entities/expand/PostExpandEntity.json
func ExampleEntitiesClient_Expand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewEntitiesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Expand(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<entity-id>",
		armsecurityinsights.EntityExpandParameters{
			EndTime:     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-26T00:00:00.000Z"); return t }()),
			ExpansionID: to.Ptr("<expansion-id>"),
			StartTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-25T00:00:00.000Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
