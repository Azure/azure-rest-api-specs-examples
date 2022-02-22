Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.3.1/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/InitiateJitNetworkAccessPolicy_example.json
func ExampleJitNetworkAccessPoliciesClient_Initiate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewJitNetworkAccessPoliciesClient("<subscription-id>",
		"<asc-location>", cred, nil)
	_, err = client.Initiate(ctx,
		"<resource-group-name>",
		"<jit-network-access-policy-name>",
		armsecurity.Enum56("initiate"),
		armsecurity.JitNetworkAccessPolicyInitiateRequest{
			Justification: to.StringPtr("<justification>"),
			VirtualMachines: []*armsecurity.JitNetworkAccessPolicyInitiateVirtualMachine{
				{
					ID: to.StringPtr("<id>"),
					Ports: []*armsecurity.JitNetworkAccessPolicyInitiatePort{
						{
							AllowedSourceAddressPrefix: to.StringPtr("<allowed-source-address-prefix>"),
							Number:                     to.Int32Ptr(3389),
						}},
				}},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
