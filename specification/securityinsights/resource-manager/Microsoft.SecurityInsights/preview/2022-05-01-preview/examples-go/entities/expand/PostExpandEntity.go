package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-05-01-preview/examples/entities/expand/PostExpandEntity.json
func ExampleEntitiesClient_Expand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewEntitiesClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Expand(ctx,
		"myRg",
		"myWorkspace",
		"e1d3d618-e11f-478b-98e3-bb381539a8e1",
		armsecurityinsights.EntityExpandParameters{
			EndTime:     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-26T00:00:00.000Z"); return t }()),
			ExpansionID: to.Ptr("a77992f3-25e9-4d01-99a4-5ff606cc410a"),
			StartTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-25T00:00:00.000Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
