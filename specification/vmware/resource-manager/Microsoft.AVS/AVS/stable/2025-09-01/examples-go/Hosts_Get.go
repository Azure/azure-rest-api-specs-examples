package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/Hosts_Get.json
func ExampleHostsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHostsClient().Get(ctx, "group1", "cloud1", "cluster1", "esx03-r52.1111111111111111111.westcentralus.prod.azure.com", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armavs.HostsClientGetResponse{
	// 	Host: &armavs.Host{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/Hosts/esx03-r52.1111111111111111111.westcentralus.prod.azure.com"),
	// 		Name: to.Ptr("esx03-r52.1111111111111111111.westcentralus.prod.azure.com"),
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/hosts"),
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 		},
	// 		SKU: &armavs.SKU{
	// 			Name: to.Ptr("av64"),
	// 		},
	// 		Properties: &armavs.GeneralHostProperties{
	// 			Kind: to.Ptr(armavs.HostKindGeneral),
	// 			DisplayName: to.Ptr("esx03-r52.1111111111111111111.westcentralus.prod.azure.com"),
	// 			MoRefID: to.Ptr("host-209"),
	// 			Fqdn: to.Ptr("esx03-r52.1111111111111111111.westcentralus.prod.azure.com"),
	// 			FaultDomain: to.Ptr("1"),
	// 		},
	// 	},
	// }
}
