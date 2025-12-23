package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/455d20a5e76d8184f7cff960501a57e1f88986b7/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-10-02-preview/examples/MachineCreate_Update.json
func ExampleMachinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMachinesClient().BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", "machine1", armcontainerservice.Machine{
		Properties: &armcontainerservice.MachineProperties{
			Hardware: &armcontainerservice.MachineHardwareProfile{
				VMSize: to.Ptr("Standard_DS1_v2"),
			},
			Kubernetes: &armcontainerservice.MachineKubernetesProfile{
				KubeletDiskType: to.Ptr(armcontainerservice.KubeletDiskTypeOS),
				MaxPods:         to.Ptr[int32](110),
				NodeLabels: map[string]*string{
					"key1": to.Ptr("val1"),
				},
				NodeTaints: []*string{
					to.Ptr("Key1=Value1:NoSchedule")},
				OrchestratorVersion: to.Ptr("1.30"),
			},
			Mode: to.Ptr(armcontainerservice.AgentPoolModeUser),
			OperatingSystem: &armcontainerservice.MachineOSProfile{
				EnableFIPS: to.Ptr(false),
				OSSKU:      to.Ptr(armcontainerservice.OSSKUUbuntu),
				OSType:     to.Ptr(armcontainerservice.OSTypeLinux),
			},
			Priority: to.Ptr(armcontainerservice.ScaleSetPrioritySpot),
			Tags: map[string]*string{
				"name1": to.Ptr("val1"),
			},
		},
		Zones: []*string{
			to.Ptr("1")},
	}, &armcontainerservice.MachinesClientBeginCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
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
	// res.Machine = armcontainerservice.Machine{
	// 	Name: to.Ptr("machine1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/agentPools/machines"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentPools/agentpool1/machines/machine1"),
	// 	Properties: &armcontainerservice.MachineProperties{
	// 		Hardware: &armcontainerservice.MachineHardwareProfile{
	// 			VMSize: to.Ptr("Standard_DS1_v2"),
	// 		},
	// 		Kubernetes: &armcontainerservice.MachineKubernetesProfile{
	// 			CurrentOrchestratorVersion: to.Ptr("1.30.6"),
	// 			KubeletDiskType: to.Ptr(armcontainerservice.KubeletDiskTypeOS),
	// 			MaxPods: to.Ptr[int32](110),
	// 			NodeLabels: map[string]*string{
	// 				"key1": to.Ptr("val1"),
	// 			},
	// 			NodeName: to.Ptr("aks-nodepool1-machine1-25481572-vm0"),
	// 			NodeTaints: []*string{
	// 				to.Ptr("Key1=Value1:NoSchedule")},
	// 				OrchestratorVersion: to.Ptr("1.30"),
	// 			},
	// 			Mode: to.Ptr(armcontainerservice.AgentPoolModeUser),
	// 			NodeImageVersion: to.Ptr("AKSUbuntu:1604:2023.03.11"),
	// 			OperatingSystem: &armcontainerservice.MachineOSProfile{
	// 				EnableFIPS: to.Ptr(false),
	// 				OSSKU: to.Ptr(armcontainerservice.OSSKUUbuntu),
	// 				OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 			},
	// 			Priority: to.Ptr(armcontainerservice.ScaleSetPrioritySpot),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			Status: &armcontainerservice.MachineStatus{
	// 				CreationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-02T12:00:00.000Z"); return t}()),
	// 				DriftAction: to.Ptr(armcontainerservice.DriftActionSynced),
	// 				VMState: to.Ptr(armcontainerservice.VMStateRunning),
	// 			},
	// 			Tags: map[string]*string{
	// 				"name1": to.Ptr("val1"),
	// 			},
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("1")},
	// 		}
}
