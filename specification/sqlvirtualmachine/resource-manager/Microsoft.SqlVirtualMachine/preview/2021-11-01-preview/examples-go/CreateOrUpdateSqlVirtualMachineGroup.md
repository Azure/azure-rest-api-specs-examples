```go
package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/CreateOrUpdateSqlVirtualMachineGroup.json
func ExampleGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewGroupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testrg",
		"testvmgroup",
		armsqlvirtualmachine.Group{
			Location: to.Ptr("northeurope"),
			Tags: map[string]*string{
				"mytag": to.Ptr("myval"),
			},
			Properties: &armsqlvirtualmachine.GroupProperties{
				SQLImageOffer: to.Ptr("SQL2016-WS2016"),
				SQLImageSKU:   to.Ptr(armsqlvirtualmachine.SQLVMGroupImageSKUEnterprise),
				WsfcDomainProfile: &armsqlvirtualmachine.WsfcDomainProfile{
					ClusterBootstrapAccount:  to.Ptr("testrpadmin"),
					ClusterOperatorAccount:   to.Ptr("testrp@testdomain.com"),
					DomainFqdn:               to.Ptr("testdomain.com"),
					OuPath:                   to.Ptr("OU=WSCluster,DC=testdomain,DC=com"),
					SQLServiceAccount:        to.Ptr("sqlservice@testdomain.com"),
					StorageAccountPrimaryKey: to.Ptr("<primary storage access key>"),
					StorageAccountURL:        to.Ptr("https://storgact.blob.core.windows.net/"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsqlvirtualmachine%2Farmsqlvirtualmachine%2Fv0.6.0/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine/README.md) on how to add the SDK to your project and authenticate.
