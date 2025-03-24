package armhybridkubernetes_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ecee919199a39cc0d864410f540aa105bf7cdb64/specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/ConnectedClustersListClusterCredentialResultHPAAD.json
func ExampleConnectedClusterClient_ListClusterUserCredential_listClusterUserCredentialCspExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridkubernetes.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedClusterClient().ListClusterUserCredential(ctx, "k8sc-rg", "testCluster", armhybridkubernetes.ListClusterUserCredentialProperties{
		AuthenticationMethod: to.Ptr(armhybridkubernetes.AuthenticationMethodAAD),
		ClientProxy:          to.Ptr(false),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CredentialResults = armhybridkubernetes.CredentialResults{
	// 	Kubeconfigs: []*armhybridkubernetes.CredentialResult{
	// 		{
	// 			Name: to.Ptr("credentialName1"),
	// 			Value: []byte("WVhCcFZtVnljMmx2YmpvZ2RqRU5DbU5zZFhOMFpYSnpPZzBLTFNCamJIVnpkR1Z5T2cwS0lDQWdJR05sY25ScFptbGpZWFJsTFdGMWRHaHZjbWwwZVMxa1lYUmhPaUJNVXpCMFRGTXhRMUpWWkVwVWFVSkVVbFpLVlZOVldrcFJNRVpWVWxNd2RFeFRNSFJEYXpGS1UxVldOR1ZyVGtSUldFVnlXakJHTTFOVlNrSmFNR3hTVlRKMGRWZHNXblphZWtwMVZtcEtWbU5ZWkV0amJsWllUVEZDU0dWclJrOVJiV1J5WTFkb2NtRlZZelZrZWtKRFVWWkdlbEpyUmtWUlZUUkxWRlpHZW1Rd1RsSlhWVkpYVlZaR1JWSllaRXRoYkdWc1NsRXdiSEZSVlRWRFdqSjBlR0ZIZEhCU2Vtd3pUVVZLUWxWVlZrZFJWVVpRVVRCR2JrOUZSazVUVld4RVVUSmtURkV3Um01U1ZVVXdWMWhDTlVOc1VtdFVNVkpTVTFkTmRtVnNhRVJsUjNoVFpXdFZNRg=="),
	// 	}},
	// }
}
