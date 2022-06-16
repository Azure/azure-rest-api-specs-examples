package armazurearcdata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/CreateOrUpdateActiveDirectoryConnector.json
func ExampleActiveDirectoryConnectorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armazurearcdata.NewActiveDirectoryConnectorsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"testrg",
		"testdataController",
		"testADConnector",
		armazurearcdata.ActiveDirectoryConnectorResource{
			Properties: &armazurearcdata.ActiveDirectoryConnectorProperties{
				Spec: &armazurearcdata.ActiveDirectoryConnectorSpec{
					ActiveDirectory: &armazurearcdata.ActiveDirectoryConnectorDomainDetails{
						DomainControllers: &armazurearcdata.ActiveDirectoryDomainControllers{
							PrimaryDomainController: &armazurearcdata.ActiveDirectoryDomainController{
								Hostname: to.Ptr("dc1.contoso.local"),
							},
							SecondaryDomainControllers: []*armazurearcdata.ActiveDirectoryDomainController{
								{
									Hostname: to.Ptr("dc2.contoso.local"),
								},
								{
									Hostname: to.Ptr("dc3.contoso.local"),
								}},
						},
						Realm:                      to.Ptr("CONTOSO.LOCAL"),
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
