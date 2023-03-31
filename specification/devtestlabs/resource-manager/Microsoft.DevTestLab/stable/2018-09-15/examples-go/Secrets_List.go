package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Secrets_List.json
func ExampleSecretsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecretsClient().NewListPager("resourceGroupName", "{labName}", "{userName}", &armdevtestlabs.SecretsClientListOptions{Expand: nil,
		Filter:  nil,
		Top:     nil,
		Orderby: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SecretList = armdevtestlabs.SecretList{
		// 	Value: []*armdevtestlabs.Secret{
		// 		{
		// 			Name: to.Ptr("secret1"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs/users/secrets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/users/{userName}/secrets/secret1"),
		// 			Properties: &armdevtestlabs.SecretProperties{
		// 				UniqueIdentifier: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("secret2"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs/users/secrets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/users/{userName}/secrets/secret2"),
		// 			Properties: &armdevtestlabs.SecretProperties{
		// 				UniqueIdentifier: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			},
		// 	}},
		// }
	}
}
