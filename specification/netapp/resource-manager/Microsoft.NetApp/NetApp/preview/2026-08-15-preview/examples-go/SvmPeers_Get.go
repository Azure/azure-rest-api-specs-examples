package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v11"
)

// Generated from example definition: 2026-08-15-preview/SvmPeers_Get.json
func ExampleSVMPeersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSVMPeersClient().Get(ctx, "contosoNetAppRg", "contoso-external-cluster", "peerA", "contoso-svm-peer-a", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.SVMPeersClientGetResponse{
	// 	SVMPeer: armnetapp.SVMPeer{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/externalClusters/contoso-external-cluster/clusterPeers/peerA/svmPeers/contoso-svm-peer-a"),
	// 		Name: to.Ptr("contoso-svm-peer-a"),
	// 		Type: to.Ptr("Microsoft.NetApp/externalClusters/clusterPeers/svmPeers"),
	// 		Properties: &armnetapp.SVMPeerProperties{
	// 			AnfSVMName: to.Ptr("anf-svm"),
	// 			ExternalSVMResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/externalClusters/contoso-external-cluster/externalSvms/contoso-svm"),
	// 			Status: to.Ptr(armnetapp.SVMPeerStatusActive),
	// 			AcceptCommand: to.Ptr("vserver peer accept -vserver contoso-svm -peer-vserver anf-svm"),
	// 			PeerTypes: []*armnetapp.SVMPeerType{
	// 				to.Ptr(armnetapp.SVMPeerTypeReplication),
	// 			},
	// 			HybridReplicationResourceIDs: []*string{
	// 				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/hybridReplications/contosoHybridReplication"),
	// 			},
	// 		},
	// 	},
	// }
}
