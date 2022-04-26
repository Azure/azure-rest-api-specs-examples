Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsqlvirtualmachine%2Farmsqlvirtualmachine%2Fv0.4.0/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine/README.md) on how to add the SDK to your project and authenticate.

```go
package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/CreateOrUpdateAvailabilityGroupListener.json
func ExampleAvailabilityGroupListenersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewAvailabilityGroupListenersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<sql-virtual-machine-group-name>",
		"<availability-group-listener-name>",
		armsqlvirtualmachine.AvailabilityGroupListener{
			Properties: &armsqlvirtualmachine.AvailabilityGroupListenerProperties{
				AvailabilityGroupName: to.Ptr("<availability-group-name>"),
				LoadBalancerConfigurations: []*armsqlvirtualmachine.LoadBalancerConfiguration{
					{
						LoadBalancerResourceID: to.Ptr("<load-balancer-resource-id>"),
						PrivateIPAddress: &armsqlvirtualmachine.PrivateIPAddress{
							IPAddress:        to.Ptr("<ipaddress>"),
							SubnetResourceID: to.Ptr("<subnet-resource-id>"),
						},
						ProbePort: to.Ptr[int32](59983),
						SQLVirtualMachineInstances: []*string{
							to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2"),
							to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm3")},
					}},
				Port: to.Ptr[int32](1433),
			},
		},
		&armsqlvirtualmachine.AvailabilityGroupListenersClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
