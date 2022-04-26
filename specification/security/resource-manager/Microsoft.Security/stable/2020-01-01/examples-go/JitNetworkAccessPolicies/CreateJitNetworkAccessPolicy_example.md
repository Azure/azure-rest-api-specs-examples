Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.6.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

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
		return
	}
	ctx := context.Background()
	client, err := armsecurity.NewJitNetworkAccessPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<asc-location>",
		"<jit-network-access-policy-name>",
		armsecurity.JitNetworkAccessPolicy{
			Kind:     to.Ptr("<kind>"),
			Location: to.Ptr("<location>"),
			Name:     to.Ptr("<name>"),
			Type:     to.Ptr("<type>"),
			ID:       to.Ptr("<id>"),
			Properties: &armsecurity.JitNetworkAccessPolicyProperties{
				ProvisioningState: to.Ptr("<provisioning-state>"),
				Requests: []*armsecurity.JitNetworkAccessRequest{
					{
						Requestor:    to.Ptr("<requestor>"),
						StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T08:06:45.5691611Z"); return t }()),
						VirtualMachines: []*armsecurity.JitNetworkAccessRequestVirtualMachine{
							{
								ID: to.Ptr("<id>"),
								Ports: []*armsecurity.JitNetworkAccessRequestPort{
									{
										AllowedSourceAddressPrefix: to.Ptr("<allowed-source-address-prefix>"),
										EndTimeUTC:                 to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T09:06:45.5691611Z"); return t }()),
										Number:                     to.Ptr[int32](3389),
										Status:                     to.Ptr(armsecurity.StatusInitiated),
										StatusReason:               to.Ptr(armsecurity.StatusReasonUserRequested),
									}},
							}},
					}},
				VirtualMachines: []*armsecurity.JitNetworkAccessPolicyVirtualMachine{
					{
						ID: to.Ptr("<id>"),
						Ports: []*armsecurity.JitNetworkAccessPortRule{
							{
								AllowedSourceAddressPrefix: to.Ptr("<allowed-source-address-prefix>"),
								MaxRequestAccessDuration:   to.Ptr("<max-request-access-duration>"),
								Number:                     to.Ptr[int32](22),
								Protocol:                   to.Ptr(armsecurity.ProtocolAll),
							},
							{
								AllowedSourceAddressPrefix: to.Ptr("<allowed-source-address-prefix>"),
								MaxRequestAccessDuration:   to.Ptr("<max-request-access-duration>"),
								Number:                     to.Ptr[int32](3389),
								Protocol:                   to.Ptr(armsecurity.ProtocolAll),
							}},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
