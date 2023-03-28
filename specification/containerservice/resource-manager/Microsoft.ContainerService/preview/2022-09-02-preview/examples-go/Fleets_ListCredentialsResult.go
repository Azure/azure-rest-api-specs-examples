package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/df863270270ad5b54fa8cce71d2c33becee0c097/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-09-02-preview/examples/Fleets_ListCredentialsResult.json
func ExampleFleetsClient_ListCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFleetsClient().ListCredentials(ctx, "rg1", "fleet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FleetCredentialResults = armcontainerservice.FleetCredentialResults{
	// 	Kubeconfigs: []*armcontainerservice.FleetCredentialResult{
	// 		{
	// 			Name: to.Ptr("credentialName1"),
	// 			Value: []byte("credentialValue1"),
	// 	}},
	// }
}
