package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsCreate.json
func ExampleComponentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewComponentsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"my-resource-group",
		"my-component",
		armapplicationinsights.Component{
			Location: to.Ptr("South Central US"),
			Kind:     to.Ptr("web"),
			Properties: &armapplicationinsights.ComponentProperties{
				ApplicationType:     to.Ptr(armapplicationinsights.ApplicationTypeWeb),
				FlowType:            to.Ptr(armapplicationinsights.FlowTypeBluefield),
				RequestSource:       to.Ptr(armapplicationinsights.RequestSourceRest),
				WorkspaceResourceID: to.Ptr("/subscriptions/subid/resourcegroups/my-resource-group/providers/microsoft.operationalinsights/workspaces/my-workspace"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
