package armpowerplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/powerplatform/resource-manager/Microsoft.PowerPlatform/preview/2020-10-30-preview/examples/createOrUpdateAccount.json
func ExampleAccountsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().CreateOrUpdate(ctx, "account", "resourceGroup", armpowerplatform.Account{
		Location: to.Ptr("East US"),
		Tags: map[string]*string{
			"Organization": to.Ptr("Administration"),
		},
		Properties: &armpowerplatform.AccountProperties{
			Description: to.Ptr("Description of the account."),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armpowerplatform.Account{
	// 	Name: to.Ptr("account"),
	// 	Type: to.Ptr("Microsoft.PowerPlatform/accounts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/resourceGroup/providers/Microsoft.PowerPlatform/accounts/account"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"Organization": to.Ptr("Administration"),
	// 	},
	// 	Properties: &armpowerplatform.AccountProperties{
	// 		Description: to.Ptr("Description of the account."),
	// 	},
	// 	SystemData: &armpowerplatform.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armpowerplatform.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armpowerplatform.CreatedByTypeUser),
	// 	},
	// }
}
