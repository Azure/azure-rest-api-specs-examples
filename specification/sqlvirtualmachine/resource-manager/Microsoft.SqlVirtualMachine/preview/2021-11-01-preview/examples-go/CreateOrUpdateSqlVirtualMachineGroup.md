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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/CreateOrUpdateSqlVirtualMachineGroup.json
func ExampleGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<sql-virtual-machine-group-name>",
		armsqlvirtualmachine.Group{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"mytag": to.Ptr("myval"),
			},
			Properties: &armsqlvirtualmachine.GroupProperties{
				SQLImageOffer: to.Ptr("<sqlimage-offer>"),
				SQLImageSKU:   to.Ptr(armsqlvirtualmachine.SQLVMGroupImageSKUEnterprise),
				WsfcDomainProfile: &armsqlvirtualmachine.WsfcDomainProfile{
					ClusterBootstrapAccount:  to.Ptr("<cluster-bootstrap-account>"),
					ClusterOperatorAccount:   to.Ptr("<cluster-operator-account>"),
					DomainFqdn:               to.Ptr("<domain-fqdn>"),
					OuPath:                   to.Ptr("<ou-path>"),
					SQLServiceAccount:        to.Ptr("<sqlservice-account>"),
					StorageAccountPrimaryKey: to.Ptr("<storage-account-primary-key>"),
					StorageAccountURL:        to.Ptr("<storage-account-url>"),
				},
			},
		},
		&armsqlvirtualmachine.GroupsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
