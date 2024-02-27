package armapicenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apicenter/armapicenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Deployments_CreateOrUpdate.json
func ExampleDeploymentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeploymentsClient().CreateOrUpdate(ctx, "contoso-resources", "contoso", "default", "echo-api", "production", armapicenter.Deployment{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Deployment = armapicenter.Deployment{
	// 	Name: to.Ptr("production"),
	// 	Type: to.Ptr("Microsoft.ApiCenter/services/apis/deployments"),
	// 	ID: to.Ptr("/subscriptions/a200340d-6b82-494d-9dbf-687ba6e33f9e/resourceGroups/contoso-resources/providers/Microsoft.ApiCenter/services/contoso/workspaces/default/deployments/production"),
	// 	SystemData: &armapicenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128Z"); return t}()),
	// 	},
	// 	Properties: &armapicenter.DeploymentProperties{
	// 		Description: to.Ptr("Public cloud production deployment."),
	// 		DefinitionID: to.Ptr("/workspaces/default/apis/echo-api/versions/2023-01-01/definitions/openapi"),
	// 		EnvironmentID: to.Ptr("/workspaces/default/environments/production"),
	// 		Server: &armapicenter.DeploymentServer{
	// 			RuntimeURI: []*string{
	// 				to.Ptr("https://api.contoso.com")},
	// 			},
	// 			State: to.Ptr(armapicenter.DeploymentStateActive),
	// 			Title: to.Ptr("Production deployment"),
	// 		},
	// 	}
}
