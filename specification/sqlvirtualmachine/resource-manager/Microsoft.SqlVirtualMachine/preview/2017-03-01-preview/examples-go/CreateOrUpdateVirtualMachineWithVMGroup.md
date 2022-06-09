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

// x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2017-03-01-preview/examples/CreateOrUpdateVirtualMachineWithVMGroup.json
func ExampleSQLVirtualMachinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsqlvirtualmachine.NewSQLVirtualMachinesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<sql-virtual-machine-name>",
		armsqlvirtualmachine.SQLVirtualMachine{
			Location: to.StringPtr("<location>"),
			Properties: &armsqlvirtualmachine.Properties{
				SQLVirtualMachineGroupResourceID: to.StringPtr("<sqlvirtual-machine-group-resource-id>"),
				VirtualMachineResourceID:         to.StringPtr("<virtual-machine-resource-id>"),
				WsfcDomainCredentials: &armsqlvirtualmachine.WsfcDomainCredentials{
					ClusterBootstrapAccountPassword: to.StringPtr("<cluster-bootstrap-account-password>"),
					ClusterOperatorAccountPassword:  to.StringPtr("<cluster-operator-account-password>"),
					SQLServiceAccountPassword:       to.StringPtr("<sqlservice-account-password>"),
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
	log.Printf("Response result: %#v\n", res.SQLVirtualMachinesClientCreateOrUpdateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsqlvirtualmachine%2Farmsqlvirtualmachine%2Fv0.2.1/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine/README.md) on how to add the SDK to your project and authenticate.
