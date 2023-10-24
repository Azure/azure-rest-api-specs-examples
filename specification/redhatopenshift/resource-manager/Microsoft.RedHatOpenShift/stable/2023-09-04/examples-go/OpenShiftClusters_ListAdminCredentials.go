package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bce66ff64f0e9edc9ea6119d00324058413e81ed/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2023-09-04/examples/OpenShiftClusters_ListAdminCredentials.json
func ExampleOpenShiftClustersClient_ListAdminCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshift.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOpenShiftClustersClient().ListAdminCredentials(ctx, "resourceGroup", "resourceName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OpenShiftClusterAdminKubeconfig = armredhatopenshift.OpenShiftClusterAdminKubeconfig{
	// 	Kubeconfig: to.Ptr("e30="),
	// }
}
