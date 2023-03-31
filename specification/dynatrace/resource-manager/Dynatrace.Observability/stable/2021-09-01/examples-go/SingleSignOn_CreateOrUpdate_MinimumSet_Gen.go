package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3751f321654db00858e70649291af5c81e94611e/specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/SingleSignOn_CreateOrUpdate_MinimumSet_Gen.json
func ExampleSingleSignOnClient_BeginCreateOrUpdate_singleSignOnCreateOrUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdynatrace.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSingleSignOnClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myMonitor", "default", armdynatrace.SingleSignOnResource{
		Properties: &armdynatrace.SingleSignOnProperties{
			AADDomains: []*string{
				to.Ptr("mpliftrdt20210811outlook.onmicrosoft.com")},
			SingleSignOnURL: to.Ptr("https://www.dynatrace.io"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SingleSignOnResource = armdynatrace.SingleSignOnResource{
	// 	Properties: &armdynatrace.SingleSignOnProperties{
	// 		AADDomains: []*string{
	// 			to.Ptr("mpliftrdt20210811outlook.onmicrosoft.com")},
	// 			SingleSignOnURL: to.Ptr("https://www.dynatrace.io/IAmSomeHash"),
	// 		},
	// 	}
}
