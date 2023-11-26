package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/Accounts_Create.json
func ExampleAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreate(ctx, "SampleResourceGroup", "Account1", armdatashare.Account{
		Location: to.Ptr("West US 2"),
		Tags: map[string]*string{
			"tag1": to.Ptr("Red"),
			"tag2": to.Ptr("White"),
		},
		Identity: &armdatashare.Identity{
			Type: to.Ptr(armdatashare.TypeSystemAssigned),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armdatashare.Account{
	// 	Name: to.Ptr("Account1"),
	// 	Type: to.Ptr("Microsoft.DataShare/accounts"),
	// 	ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1"),
	// 	SystemData: &armdatashare.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdatashare.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armdatashare.LastModifiedByTypeUser),
	// 	},
	// 	Location: to.Ptr("West US 2"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("Red"),
	// 		"tag2": to.Ptr("White"),
	// 	},
	// 	Identity: &armdatashare.Identity{
	// 		Type: to.Ptr(armdatashare.TypeSystemAssigned),
	// 	},
	// 	Properties: &armdatashare.AccountProperties{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-14T04:47:52.961Z"); return t}()),
	// 		UserEmail: to.Ptr("johnsmith@microsoft.com"),
	// 		UserName: to.Ptr("John Smith"),
	// 	},
	// }
}
