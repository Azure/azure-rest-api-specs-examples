package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9cc7633f842575274f715cc02e37c5769ac2742d/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/SnapshotsGet.json
func ExampleSnapshotsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSnapshotsClient().Get(ctx, "rg1", "snapshot1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Snapshot = armcontainerservice.Snapshot{
	// 	Name: to.Ptr("snapshot1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/Snapshots"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/snapshots/snapshot1"),
	// 	SystemData: &armcontainerservice.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-09T20:13:23.298Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("val1"),
	// 		"key2": to.Ptr("val2"),
	// 	},
	// 	Properties: &armcontainerservice.SnapshotProperties{
	// 		CreationData: &armcontainerservice.CreationData{
	// 			SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster1/agentPools/pool0"),
	// 		},
	// 		EnableFIPS: to.Ptr(false),
	// 		KubernetesVersion: to.Ptr("1.20.5"),
	// 		NodeImageVersion: to.Ptr("AKSUbuntu-1804gen2containerd-2021.09.11"),
	// 		OSSKU: to.Ptr(armcontainerservice.OSSKUUbuntu),
	// 		OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 		SnapshotType: to.Ptr(armcontainerservice.SnapshotTypeNodePool),
	// 		VMSize: to.Ptr("Standard_D2s_v3"),
	// 	},
	// }
}
