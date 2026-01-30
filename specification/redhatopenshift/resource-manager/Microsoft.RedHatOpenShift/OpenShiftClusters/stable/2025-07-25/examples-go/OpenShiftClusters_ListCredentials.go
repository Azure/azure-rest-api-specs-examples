package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6e34caed36815fc876c8e8c0371db76f809e52e8/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/OpenShiftClusters/stable/2025-07-25/examples/OpenShiftClusters_ListCredentials.json
func ExampleOpenShiftClustersClient_ListCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshift.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOpenShiftClustersClient().ListCredentials(ctx, "resourceGroup", "resourceName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OpenShiftClusterCredentials = armredhatopenshift.OpenShiftClusterCredentials{
	// 	KubeadminPassword: to.Ptr("password"),
	// 	KubeadminUsername: to.Ptr("kubeadmin"),
	// }
}
