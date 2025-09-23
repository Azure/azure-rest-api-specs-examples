package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1d2097f1ed03e8a61eed4fe63602a641bedd77ae/specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/AppResiliency_List.json
func ExampleAppResiliencyClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAppResiliencyClient().NewListPager("rg", "testcontainerApp0", nil)
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
		// page.AppResiliencyCollection = armappcontainers.AppResiliencyCollection{
		// 	Value: []*armappcontainers.AppResiliency{
		// 		{
		// 			Name: to.Ptr("resiliency-policy-1"),
		// 			Type: to.Ptr("Microsoft.App/containerApps/resiliencyPolicies"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/testcontainerApp0/resiliencyPolicies/resiliency-policy-1"),
		// 			Properties: &armappcontainers.AppResiliencyProperties{
		// 				CircuitBreakerPolicy: &armappcontainers.CircuitBreakerPolicy{
		// 					ConsecutiveErrors: to.Ptr[int32](5),
		// 					IntervalInSeconds: to.Ptr[int32](10),
		// 					MaxEjectionPercent: to.Ptr[int32](50),
		// 				},
		// 				HTTPConnectionPool: &armappcontainers.HTTPConnectionPool{
		// 					HTTP1MaxPendingRequests: to.Ptr[int32](1024),
		// 					HTTP2MaxRequests: to.Ptr[int32](1024),
		// 				},
		// 				HTTPRetryPolicy: &armappcontainers.HTTPRetryPolicy{
		// 					Matches: &armappcontainers.HTTPRetryPolicyMatches{
		// 						Errors: []*string{
		// 							to.Ptr("5xx"),
		// 							to.Ptr("connect-failure"),
		// 							to.Ptr("reset"),
		// 							to.Ptr("retriable-headers"),
		// 							to.Ptr("retriable-status-codes")},
		// 							Headers: []*armappcontainers.HeaderMatch{
		// 								{
		// 									Header: to.Ptr("X-Content-Type"),
		// 									Match: &armappcontainers.HeaderMatchMatch{
		// 										PrefixMatch: to.Ptr("GOATS"),
		// 									},
		// 							}},
		// 							HTTPStatusCodes: []*int32{
		// 								to.Ptr[int32](502),
		// 								to.Ptr[int32](503)},
		// 							},
		// 							MaxRetries: to.Ptr[int32](5),
		// 							RetryBackOff: &armappcontainers.HTTPRetryPolicyRetryBackOff{
		// 								InitialDelayInMilliseconds: to.Ptr[int64](1000),
		// 								MaxIntervalInMilliseconds: to.Ptr[int64](10000),
		// 							},
		// 						},
		// 						TCPConnectionPool: &armappcontainers.TCPConnectionPool{
		// 							MaxConnections: to.Ptr[int32](100),
		// 						},
		// 						TCPRetryPolicy: &armappcontainers.TCPRetryPolicy{
		// 							MaxConnectAttempts: to.Ptr[int32](3),
		// 						},
		// 						TimeoutPolicy: &armappcontainers.TimeoutPolicy{
		// 							ConnectionTimeoutInSeconds: to.Ptr[int32](5),
		// 							ResponseTimeoutInSeconds: to.Ptr[int32](15),
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
