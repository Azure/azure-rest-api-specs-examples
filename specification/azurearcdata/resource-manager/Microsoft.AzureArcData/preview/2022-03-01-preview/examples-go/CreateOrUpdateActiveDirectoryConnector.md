Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazurearcdata%2Farmazurearcdata%2Fv0.4.0/sdk/resourcemanager/azurearcdata/armazurearcdata/README.md) on how to add the SDK to your project and authenticate.

```go
package armazurearcdata_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/CreateOrUpdateActiveDirectoryConnector.json
func ExampleActiveDirectoryConnectorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armazurearcdata.NewActiveDirectoryConnectorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<data-controller-name>",
		"<active-directory-connector-name>",
		armazurearcdata.ActiveDirectoryConnectorResource{
			Properties: &armazurearcdata.ActiveDirectoryConnectorProperties{
				Spec: &armazurearcdata.ActiveDirectoryConnectorSpec{
					ActiveDirectory: &armazurearcdata.ActiveDirectoryConnectorDomainDetails{
						DomainControllers: &armazurearcdata.ActiveDirectoryDomainControllers{
							PrimaryDomainController: &armazurearcdata.ActiveDirectoryDomainController{
								Hostname: to.Ptr("<hostname>"),
							},
							SecondaryDomainControllers: []*armazurearcdata.ActiveDirectoryDomainController{
								{
									Hostname: to.Ptr("<hostname>"),
								},
								{
									Hostname: to.Ptr("<hostname>"),
								}},
						},
						Realm:                      to.Ptr("<realm>"),
						ServiceAccountProvisioning: to.Ptr(armazurearcdata.AccountProvisioningModeManual),
					},
					DNS: &armazurearcdata.ActiveDirectoryConnectorDNSDetails{
						NameserverIPAddresses: []*string{
							to.Ptr("11.11.111.111"),
							to.Ptr("22.22.222.222")},
						PreferK8SDNSForPtrLookups: to.Ptr(false),
						Replicas:                  to.Ptr[int64](1),
					},
				},
			},
		},
		&armazurearcdata.ActiveDirectoryConnectorsClientBeginCreateOptions{ResumeToken: ""})
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
