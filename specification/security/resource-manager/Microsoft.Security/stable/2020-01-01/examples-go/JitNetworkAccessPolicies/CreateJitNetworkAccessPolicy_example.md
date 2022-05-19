Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.7.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/CreateJitNetworkAccessPolicy_example.json
func ExampleJitNetworkAccessPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewJitNetworkAccessPoliciesClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"myRg1",
		"westeurope",
		"default",
		armsecurity.JitNetworkAccessPolicy{
			Kind:     to.Ptr("Basic"),
			Location: to.Ptr("westeurope"),
			Name:     to.Ptr("default"),
			Type:     to.Ptr("Microsoft.Security/locations/jitNetworkAccessPolicies"),
			ID:       to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Security/locations/westeurope/jitNetworkAccessPolicies/default"),
			Properties: &armsecurity.JitNetworkAccessPolicyProperties{
				ProvisioningState: to.Ptr("Succeeded"),
				Requests: []*armsecurity.JitNetworkAccessRequest{
					{
						Requestor:    to.Ptr("barbara@contoso.com"),
						StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T08:06:45.5691611Z"); return t }()),
						VirtualMachines: []*armsecurity.JitNetworkAccessRequestVirtualMachine{
							{
								ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
								Ports: []*armsecurity.JitNetworkAccessRequestPort{
									{
										AllowedSourceAddressPrefix: to.Ptr("192.127.0.2"),
										EndTimeUTC:                 to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T09:06:45.5691611Z"); return t }()),
										Number:                     to.Ptr[int32](3389),
										Status:                     to.Ptr(armsecurity.StatusInitiated),
										StatusReason:               to.Ptr(armsecurity.StatusReasonUserRequested),
									}},
							}},
					}},
				VirtualMachines: []*armsecurity.JitNetworkAccessPolicyVirtualMachine{
					{
						ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
						Ports: []*armsecurity.JitNetworkAccessPortRule{
							{
								AllowedSourceAddressPrefix: to.Ptr("*"),
								MaxRequestAccessDuration:   to.Ptr("PT3H"),
								Number:                     to.Ptr[int32](22),
								Protocol:                   to.Ptr(armsecurity.ProtocolAll),
							},
							{
								AllowedSourceAddressPrefix: to.Ptr("*"),
								MaxRequestAccessDuration:   to.Ptr("PT3H"),
								Number:                     to.Ptr[int32](3389),
								Protocol:                   to.Ptr(armsecurity.ProtocolAll),
							}},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
