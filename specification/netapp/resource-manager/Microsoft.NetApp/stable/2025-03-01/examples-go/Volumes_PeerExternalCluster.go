package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/Volumes_PeerExternalCluster.json
func ExampleVolumesClient_BeginPeerExternalCluster() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginPeerExternalCluster(ctx, "myRG", "account1", "pool1", "volume1", armnetapp.PeerClusterForVolumeMigrationRequest{
		PeerIPAddresses: []*string{
			to.Ptr("0.0.0.1"),
			to.Ptr("0.0.0.2"),
			to.Ptr("0.0.0.3"),
			to.Ptr("0.0.0.4"),
			to.Ptr("0.0.0.5"),
			to.Ptr("0.0.0.6")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClusterPeerCommandResponse = armnetapp.ClusterPeerCommandResponse{
	// 	PeerAcceptCommand: to.Ptr("cluster peer create -ipspace replication -encryption-protocol-proposed tls-psk -passphrase passphraseString -peer-addrs 1.1.1.1,1.1.1.2,1.1.1.3,1.1.1.4,1.1.1.5,1.1.1.6"),
	// }
}
