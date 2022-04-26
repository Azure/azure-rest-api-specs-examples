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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/CreateOrUpdateVirtualMachineWithVMGroup.json
func ExampleSQLVirtualMachinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<sql-virtual-machine-name>",
		armsqlvirtualmachine.SQLVirtualMachine{
			Location: to.Ptr("<location>"),
			Properties: &armsqlvirtualmachine.Properties{
				SQLVirtualMachineGroupResourceID: to.Ptr("<sqlvirtual-machine-group-resource-id>"),
				VirtualMachineResourceID:         to.Ptr("<virtual-machine-resource-id>"),
				WsfcDomainCredentials: &armsqlvirtualmachine.WsfcDomainCredentials{
					ClusterBootstrapAccountPassword: to.Ptr("<cluster-bootstrap-account-password>"),
					ClusterOperatorAccountPassword:  to.Ptr("<cluster-operator-account-password>"),
					SQLServiceAccountPassword:       to.Ptr("<sqlservice-account-password>"),
				},
			},
		},
		&armsqlvirtualmachine.SQLVirtualMachinesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
