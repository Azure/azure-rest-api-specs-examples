package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/DaprComponentResiliencyPolicies_List.json
func ExampleDaprComponentResiliencyPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDaprComponentResiliencyPoliciesClient().NewListPager("examplerg", "myenvironment", "mydaprcomponent", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DaprComponentResiliencyPoliciesCollection = armappcontainers.DaprComponentResiliencyPoliciesCollection{
		// 	Value: []*armappcontainers.DaprComponentResiliencyPolicy{
		// 		{
		// 			Name: to.Ptr("something"),
		// 			Type: to.Ptr("Microsoft.App/managedEnvironments/daprComponents/resiliencyPolicies"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/daprComponents/mydaprcomponent/resiliencyPolicies/myresiliencypolicy"),
		// 			Properties: &armappcontainers.DaprComponentResiliencyPolicyProperties{
		// 				InboundPolicy: &armappcontainers.DaprComponentResiliencyPolicyConfiguration{
		// 					CircuitBreakerPolicy: &armappcontainers.DaprComponentResiliencyPolicyCircuitBreakerPolicyConfiguration{
		// 						ConsecutiveErrors: to.Ptr[int32](5),
		// 						TimeoutInSeconds: to.Ptr[int32](10),
		// 					},
		// 					HTTPRetryPolicy: &armappcontainers.DaprComponentResiliencyPolicyHTTPRetryPolicyConfiguration{
		// 						MaxRetries: to.Ptr[int32](15),
		// 						RetryBackOff: &armappcontainers.DaprComponentResiliencyPolicyHTTPRetryBackOffConfiguration{
		// 							InitialDelayInMilliseconds: to.Ptr[int32](2000),
		// 							MaxIntervalInMilliseconds: to.Ptr[int32](5500),
		// 						},
		// 					},
		// 					TimeoutPolicy: &armappcontainers.DaprComponentResiliencyPolicyTimeoutPolicyConfiguration{
		// 						ResponseTimeoutInSeconds: to.Ptr[int32](30),
		// 					},
		// 				},
		// 				OutboundPolicy: &armappcontainers.DaprComponentResiliencyPolicyConfiguration{
		// 					CircuitBreakerPolicy: &armappcontainers.DaprComponentResiliencyPolicyCircuitBreakerPolicyConfiguration{
		// 						ConsecutiveErrors: to.Ptr[int32](3),
		// 						IntervalInSeconds: to.Ptr[int32](60),
		// 						TimeoutInSeconds: to.Ptr[int32](20),
		// 					},
		// 					HTTPRetryPolicy: &armappcontainers.DaprComponentResiliencyPolicyHTTPRetryPolicyConfiguration{
		// 						MaxRetries: to.Ptr[int32](5),
		// 						RetryBackOff: &armappcontainers.DaprComponentResiliencyPolicyHTTPRetryBackOffConfiguration{
		// 							InitialDelayInMilliseconds: to.Ptr[int32](100),
		// 							MaxIntervalInMilliseconds: to.Ptr[int32](30000),
		// 						},
		// 					},
		// 					TimeoutPolicy: &armappcontainers.DaprComponentResiliencyPolicyTimeoutPolicyConfiguration{
		// 						ResponseTimeoutInSeconds: to.Ptr[int32](12),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
