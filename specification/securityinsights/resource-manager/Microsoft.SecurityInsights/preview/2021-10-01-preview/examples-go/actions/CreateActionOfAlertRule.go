package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/actions/CreateActionOfAlertRule.json
func ExampleActionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewActionsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<rule-id>",
		"<action-id>",
		armsecurityinsights.ActionRequest{
			Etag: to.StringPtr("<etag>"),
			Properties: &armsecurityinsights.ActionRequestProperties{
				LogicAppResourceID: to.StringPtr("<logic-app-resource-id>"),
				TriggerURI:         to.StringPtr("<trigger-uri>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ActionsClientCreateOrUpdateResult)
}
