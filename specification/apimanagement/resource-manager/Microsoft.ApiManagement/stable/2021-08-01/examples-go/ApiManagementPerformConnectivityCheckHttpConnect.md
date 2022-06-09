```go
package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementPerformConnectivityCheckHttpConnect.json
func ExampleClient_BeginPerformConnectivityCheckAsync() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPerformConnectivityCheckAsync(ctx,
		"rg1",
		"apimService1",
		armapimanagement.ConnectivityCheckRequest{
			Destination: &armapimanagement.ConnectivityCheckRequestDestination{
				Address: to.Ptr("https://microsoft.com"),
				Port:    to.Ptr[int64](3306),
			},
			ProtocolConfiguration: &armapimanagement.ConnectivityCheckRequestProtocolConfiguration{
				HTTPConfiguration: &armapimanagement.ConnectivityCheckRequestProtocolConfigurationHTTPConfiguration{
					Method: to.Ptr(armapimanagement.MethodGET),
					Headers: []*armapimanagement.HTTPHeader{
						{
							Name:  to.Ptr("Authorization"),
							Value: to.Ptr("Bearer myPreciousToken"),
						}},
					ValidStatusCodes: []*int64{
						to.Ptr[int64](200),
						to.Ptr[int64](204)},
				},
			},
			Source: &armapimanagement.ConnectivityCheckRequestSource{
				Region: to.Ptr("northeurope"),
			},
			Protocol: to.Ptr(armapimanagement.ConnectivityCheckProtocolHTTPS),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapimanagement%2Farmapimanagement%2Fv1.0.0/sdk/resourcemanager/apimanagement/armapimanagement/README.md) on how to add the SDK to your project and authenticate.
