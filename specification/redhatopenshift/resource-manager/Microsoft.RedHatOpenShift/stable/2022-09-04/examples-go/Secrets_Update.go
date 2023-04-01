package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d89e212da2cc34206fe711f44dfcb6f8fdece2a1/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2022-09-04/examples/Secrets_Update.json
func ExampleSecretsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshift.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecretsClient().Update(ctx, "resourceGroup", "resourceName", "childResourceName", armredhatopenshift.SecretUpdate{
		Properties: &armredhatopenshift.SecretProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Secret = armredhatopenshift.Secret{
	// 	Name: to.Ptr("mySecret"),
	// 	Type: to.Ptr("Microsoft.RedHatOpenShift/OpenShiftClusters/Secrets"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/Microsoft.RedHatOpenShift/OpenShiftClusters/resourceName/secret/mySecret"),
	// 	Properties: &armredhatopenshift.SecretProperties{
	// 	},
	// }
}
