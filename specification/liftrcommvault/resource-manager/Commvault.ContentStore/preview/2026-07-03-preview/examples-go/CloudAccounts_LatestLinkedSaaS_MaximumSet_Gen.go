package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-07-03-preview/CloudAccounts_LatestLinkedSaaS_MaximumSet_Gen.json
func ExampleCloudAccountsClient_LatestLinkedSaaS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCloudAccountsClient().LatestLinkedSaaS(ctx, "rg-commvault", "contoso-cloud-account", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcommvaultcontentstore.CloudAccountsClientLatestLinkedSaaSResponse{
	// 	LatestLinkedSaaSResponse: armcommvaultcontentstore.LatestLinkedSaaSResponse{
	// 		SaaSResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-commvault/providers/Microsoft.SaaS/resources/commvault-saas"),
	// 		IsHiddenSaaS: to.Ptr(false),
	// 	},
	// }
}
