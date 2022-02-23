Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.3.1/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/CreateJitNetworkAccessPolicy_example.json
func ExampleJitNetworkAccessPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewJitNetworkAccessPoliciesClient("<subscription-id>",
		"<asc-location>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<jit-network-access-policy-name>",
		armsecurity.JitNetworkAccessPolicy{
			Kind:     to.StringPtr("<kind>"),
			Location: to.StringPtr("<location>"),
			Name:     to.StringPtr("<name>"),
			Type:     to.StringPtr("<type>"),
			ID:       to.StringPtr("<id>"),
			Properties: &armsecurity.JitNetworkAccessPolicyProperties{
				ProvisioningState: to.StringPtr("<provisioning-state>"),
				Requests: []*armsecurity.JitNetworkAccessRequest{
					{
						Requestor:    to.StringPtr("<requestor>"),
						StartTimeUTC: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T08:06:45.5691611Z"); return t }()),
						VirtualMachines: []*armsecurity.JitNetworkAccessRequestVirtualMachine{
							{
								ID: to.StringPtr("<id>"),
								Ports: []*armsecurity.JitNetworkAccessRequestPort{
									{
										AllowedSourceAddressPrefix: to.StringPtr("<allowed-source-address-prefix>"),
										EndTimeUTC:                 to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T09:06:45.5691611Z"); return t }()),
										Number:                     to.Int32Ptr(3389),
										Status:                     armsecurity.Status("Initiated").ToPtr(),
										StatusReason:               armsecurity.StatusReason("UserRequested").ToPtr(),
									}},
							}},
					}},
				VirtualMachines: []*armsecurity.JitNetworkAccessPolicyVirtualMachine{
					{
						ID: to.StringPtr("<id>"),
						Ports: []*armsecurity.JitNetworkAccessPortRule{
							{
								AllowedSourceAddressPrefix: to.StringPtr("<allowed-source-address-prefix>"),
								MaxRequestAccessDuration:   to.StringPtr("<max-request-access-duration>"),
								Number:                     to.Int32Ptr(22),
								Protocol:                   armsecurity.Protocol("*").ToPtr(),
							},
							{
								AllowedSourceAddressPrefix: to.StringPtr("<allowed-source-address-prefix>"),
								MaxRequestAccessDuration:   to.StringPtr("<max-request-access-duration>"),
								Number:                     to.Int32Ptr(3389),
								Protocol:                   armsecurity.Protocol("*").ToPtr(),
							}},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.JitNetworkAccessPoliciesClientCreateOrUpdateResult)
}
```
