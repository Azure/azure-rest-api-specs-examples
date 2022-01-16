Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsolutions%2Farmmanagedapplications%2Fv0.2.0/sdk/resourcemanager/solutions/armmanagedapplications/README.md) on how to add the SDK to your project and authenticate.

```go
package armmanagedapplications_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications"
)

// x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/createOrUpdateJitRequest.json
func ExampleJitRequestsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagedapplications.NewJitRequestsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<jit-request-name>",
		armmanagedapplications.JitRequestDefinition{
			Properties: &armmanagedapplications.JitRequestProperties{
				ApplicationResourceID: to.StringPtr("<application-resource-id>"),
				JitAuthorizationPolicies: []*armmanagedapplications.JitAuthorizationPolicies{
					{
						PrincipalID:      to.StringPtr("<principal-id>"),
						RoleDefinitionID: to.StringPtr("<role-definition-id>"),
					}},
				JitSchedulingPolicy: &armmanagedapplications.JitSchedulingPolicy{
					Type:      armmanagedapplications.JitSchedulingType("Once").ToPtr(),
					Duration:  to.StringPtr("<duration>"),
					StartTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-22T05:48:30.6661804Z"); return t }()),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.JitRequestsClientCreateOrUpdateResult)
}
```
