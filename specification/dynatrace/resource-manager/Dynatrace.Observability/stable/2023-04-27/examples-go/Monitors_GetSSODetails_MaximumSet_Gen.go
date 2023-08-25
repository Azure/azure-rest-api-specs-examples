package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_GetSSODetails_MaximumSet_Gen.json
func ExampleMonitorsClient_GetSSODetails_monitorsGetSsoDetailsMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdynatrace.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMonitorsClient().GetSSODetails(ctx, "myResourceGroup", "myMonitor", &armdynatrace.MonitorsClientGetSSODetailsOptions{Request: &armdynatrace.SSODetailsRequest{
		UserPrincipal: to.Ptr("alice@microsoft.com"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SSODetailsResponse = armdynatrace.SSODetailsResponse{
	// 	AADDomains: []*string{
	// 		to.Ptr("mpliftrdt20210811outlook.onmicrosoft.com")},
	// 		AdminUsers: []*string{
	// 			to.Ptr("alice@microsoft.com")},
	// 			IsSsoEnabled: to.Ptr(armdynatrace.SSOStatusEnabled),
	// 			MetadataURL: to.Ptr("https://someuri"),
	// 			SingleSignOnURL: to.Ptr("https://www.dynatrace.io/IAmSomeHash"),
	// 		}
}
