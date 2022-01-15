Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsqlvirtualmachine%2Farmsqlvirtualmachine%2Fv0.2.0/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2017-03-01-preview/examples/CreateOrUpdateSqlVirtualMachineGroup.json
func ExampleGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsqlvirtualmachine.NewGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<sql-virtual-machine-group-name>",
		armsqlvirtualmachine.Group{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"mytag": to.StringPtr("myval"),
			},
			Properties: &armsqlvirtualmachine.GroupProperties{
				SQLImageOffer: to.StringPtr("<sqlimage-offer>"),
				SQLImageSKU:   armsqlvirtualmachine.SQLVMGroupImageSKU("Enterprise").ToPtr(),
				WsfcDomainProfile: &armsqlvirtualmachine.WsfcDomainProfile{
					ClusterBootstrapAccount:  to.StringPtr("<cluster-bootstrap-account>"),
					ClusterOperatorAccount:   to.StringPtr("<cluster-operator-account>"),
					DomainFqdn:               to.StringPtr("<domain-fqdn>"),
					OuPath:                   to.StringPtr("<ou-path>"),
					SQLServiceAccount:        to.StringPtr("<sqlservice-account>"),
					StorageAccountPrimaryKey: to.StringPtr("<storage-account-primary-key>"),
					StorageAccountURL:        to.StringPtr("<storage-account-url>"),
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
	log.Printf("Response result: %#v\n", res.GroupsClientCreateOrUpdateResult)
}
```
