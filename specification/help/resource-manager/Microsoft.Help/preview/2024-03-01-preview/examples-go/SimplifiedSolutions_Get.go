package armselfhelp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/selfhelp/armselfhelp/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c3cc9abe085093ba880ee3eeb792edb4fa789553/specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/SimplifiedSolutions_Get.json
func ExampleSimplifiedSolutionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSimplifiedSolutionsClient().Get(ctx, "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp", "simplifiedSolutionsResourceName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SimplifiedSolutionsResource = armselfhelp.SimplifiedSolutionsResource{
	// 	Name: to.Ptr("simplifiedSolutionsResourceName1"),
	// 	Type: to.Ptr("Microsoft.Help/simplifiedSolutions"),
	// 	ID: to.Ptr("/subscriptions/mySubscription/resourceGroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp/providers/Microsoft.Help/solutions/SolutionResource1"),
	// 	Properties: &armselfhelp.SimplifiedSolutionsResourceProperties{
	// 		Content: to.Ptr("<p>Sample content</p>"),
	// 		ProvisioningState: to.Ptr(armselfhelp.SolutionProvisioningStateSucceeded),
	// 		SolutionID: to.Ptr("sampleSolutionId1"),
	// 		Title: to.Ptr("RBAC Authentication Common Solutions"),
	// 	},
	// }
}
