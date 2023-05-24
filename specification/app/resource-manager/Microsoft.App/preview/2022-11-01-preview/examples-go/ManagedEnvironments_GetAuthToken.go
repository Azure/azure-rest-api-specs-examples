package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a7af0df86022e5e6cc6e8f40ca1981c4557a4bc/specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/ManagedEnvironments_GetAuthToken.json
func ExampleManagedEnvironmentsClient_GetAuthToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedEnvironmentsClient().GetAuthToken(ctx, "rg", "testenv", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EnvironmentAuthToken = armappcontainers.EnvironmentAuthToken{
	// 	Name: to.Ptr("testenv"),
	// 	Type: to.Ptr("Microsoft.App/environments/accesstoken"),
	// 	ID: to.Ptr("/subscriptions/651f8027-33e8-4ec4-97b4-f6e9f3dc8744/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/testenv"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappcontainers.EnvironmentAuthTokenProperties{
	// 		Expires: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-14T19:22:50.3080223Z"); return t}()),
	// 		Token: to.Ptr("testToken"),
	// 	},
	// }
}
