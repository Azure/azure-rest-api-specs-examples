package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/29bc17223449e7a865a6f38552eacfd213d812fd/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-02-02-preview/examples/ManagedClustersListClusterCredentialResult.json
func ExampleManagedClustersClient_ListClusterMonitoringUserCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedClustersClient().ListClusterMonitoringUserCredentials(ctx, "rg1", "clustername1", &armcontainerservice.ManagedClustersClientListClusterMonitoringUserCredentialsOptions{ServerFqdn: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CredentialResults = armcontainerservice.CredentialResults{
	// 	Kubeconfigs: []*armcontainerservice.CredentialResult{
	// 		{
	// 			Name: to.Ptr("credentialName1"),
	// 			Value: []byte("Y3JlZGVudGlhbFZhbHVlMQ=="),
	// 	}},
	// }
}
