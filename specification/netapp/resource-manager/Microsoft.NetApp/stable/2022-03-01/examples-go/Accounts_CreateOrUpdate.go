package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-03-01/examples/Accounts_CreateOrUpdate.json
func ExampleAccountsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetapp.NewAccountsClient("D633CC2E-722B-4AE1-B636-BBD9E4C60ED9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myRG",
		"account1",
		armnetapp.Account{
			Location: to.Ptr("eastus"),
			Properties: &armnetapp.AccountProperties{
				ActiveDirectories: []*armnetapp.ActiveDirectory{
					{
						AesEncryption:      to.Ptr(true),
						DNS:                to.Ptr("10.10.10.3, 10.10.10.4"),
						Domain:             to.Ptr("10.10.10.3"),
						LdapOverTLS:        to.Ptr(false),
						LdapSigning:        to.Ptr(false),
						OrganizationalUnit: to.Ptr("OU=Engineering"),
						Password:           to.Ptr("ad_password"),
						Site:               to.Ptr("SiteName"),
						SmbServerName:      to.Ptr("SMBServer"),
						Username:           to.Ptr("ad_user_name"),
					}},
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
