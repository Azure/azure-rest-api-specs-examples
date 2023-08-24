package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2023-04-01/examples/Secrets_List.json
func ExampleSecretsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshift.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecretsClient().NewListPager("resourceGroup", "resourceName", nil)
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
		// page.SecretList = armredhatopenshift.SecretList{
		// 	Value: []*armredhatopenshift.Secret{
		// 		{
		// 			Name: to.Ptr("mySecret"),
		// 			Type: to.Ptr("Microsoft.RedHatOpenShift/OpenShiftClusters/Secrets"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/Microsoft.RedHatOpenShift/OpenShiftClusters/resourceName/secret/mySecret"),
		// 			Properties: &armredhatopenshift.SecretProperties{
		// 			},
		// 	}},
		// }
	}
}
