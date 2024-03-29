package armhybridkubernetes_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/stable/2021-10-01/examples/ConnectedClustersListClusterCredentialResultHPAAD.json
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
	// 			Value: []byte("YXBpVmVyc2lvbjogdjENCmNsdXN0ZXJzOg0KLSBjbHVzdGVyOg0KICAgIGNlcnRpZmljYXRlLWF1dGhvcml0eS1kYXRhOiBMUzB0TFMxQ1JVZEpUaUJEUlZKVVNVWkpRMEZVUlMwdExTMHRDazFKU1VWNGVrTkRRWEVyWjBGM1NVSkJaMGxSVTJ0dVdsWnZaekp1VmpKVmNYZEtjblZYTTFCSGVrRk9RbWRyY1docmFVYzVkekJDUVZGelJrRkVRVTRLVFZGemQwTlJXVVJXVVZGRVJYZEthbGVsSlEwbHFRVTVDWjJ0eGFHdHBSemwzTUVKQlVVVkdRVUZQUTBGbk9FRk5TVWxEUTJkTFEwRm5SVUUwV1hCNUNsUmtUMVJSU1dNdmVsaERlR3hTZWtVMF"),
	// 	}},
	// }
}
