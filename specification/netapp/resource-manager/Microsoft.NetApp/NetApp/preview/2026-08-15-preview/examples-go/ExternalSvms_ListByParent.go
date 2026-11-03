package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v11"
)

// Generated from example definition: 2026-08-15-preview/ExternalSvms_ListByParent.json
func ExampleExternalSVMsClient_NewListByParentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExternalSVMsClient().NewListByParentPager("contosoNetAppRg", "contoso-external-cluster", nil)
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
		// page = armnetapp.ExternalSVMsClientListByParentResponse{
		// 	ExternalSvmListResult: armnetapp.ExternalSvmListResult{
		// 		Value: []*armnetapp.ExternalSVM{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/externalClusters/contoso-external-cluster/externalSvms/contoso-svm"),
		// 				Name: to.Ptr("contoso-svm"),
		// 				Type: to.Ptr("Microsoft.NetApp/externalClusters/externalSvms"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armnetapp.ExternalSVMProperties{
		// 					ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
		// 					OntapSVMName: to.Ptr("contoso-svm"),
		// 					SVMPeerResourceIDs: []*string{
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/externalClusters/contoso-external-cluster/clusterPeers/peerA/svmPeers/svm-peerA"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/externalClusters/contoso-external-cluster/externalSvms/svm2"),
		// 				Name: to.Ptr("svm2"),
		// 				Type: to.Ptr("Microsoft.NetApp/externalClusters/externalSvms"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armnetapp.ExternalSVMProperties{
		// 					ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
		// 					OntapSVMName: to.Ptr("svm2"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
