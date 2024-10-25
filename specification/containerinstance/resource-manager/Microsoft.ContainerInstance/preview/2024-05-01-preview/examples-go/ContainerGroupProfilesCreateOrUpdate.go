package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-05-01-preview/examples/ContainerGroupProfilesCreateOrUpdate.json
func ExampleContainerGroupProfilesClient_CreateOrUpdate_containerGroupProfilesCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerGroupProfilesClient().CreateOrUpdate(ctx, "demo", "demo1", armcontainerinstance.ContainerGroupProfile{
		Properties: &armcontainerinstance.ContainerGroupProfilePropertiesProperties{
			Containers: []*armcontainerinstance.Container{
				{
					Name: to.Ptr("demo1"),
					Properties: &armcontainerinstance.ContainerProperties{
						Command:              []*string{},
						EnvironmentVariables: []*armcontainerinstance.EnvironmentVariable{},
						Image:                to.Ptr("nginx"),
						Ports: []*armcontainerinstance.ContainerPort{
							{
								Port: to.Ptr[int32](80),
							}},
						Resources: &armcontainerinstance.ResourceRequirements{
							Requests: &armcontainerinstance.ResourceRequests{
								CPU: to.Ptr[float64](1),
								Gpu: &armcontainerinstance.GpuResource{
									Count: to.Ptr[int32](1),
									SKU:   to.Ptr(armcontainerinstance.GpuSKUK80),
								},
								MemoryInGB: to.Ptr[float64](1.5),
							},
						},
						VolumeMounts: []*armcontainerinstance.VolumeMount{
							{
								Name:      to.Ptr("volume1"),
								MountPath: to.Ptr("/mnt/volume1"),
								ReadOnly:  to.Ptr(false),
							},
							{
								Name:      to.Ptr("volume2"),
								MountPath: to.Ptr("/mnt/volume2"),
								ReadOnly:  to.Ptr(false),
							},
							{
								Name:      to.Ptr("volume3"),
								MountPath: to.Ptr("/mnt/volume3"),
								ReadOnly:  to.Ptr(true),
							}},
					},
				}},
			Diagnostics: &armcontainerinstance.ContainerGroupDiagnostics{
				LogAnalytics: &armcontainerinstance.LogAnalytics{
					LogType: to.Ptr(armcontainerinstance.LogAnalyticsLogTypeContainerInsights),
					Metadata: map[string]*string{
						"pod-uuid": to.Ptr("test-metadata-value"),
					},
					WorkspaceID:         to.Ptr("workspaceid"),
					WorkspaceKey:        to.Ptr("workspaceKey"),
					WorkspaceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/microsoft.operationalinsights/workspaces/workspace"),
				},
			},
			ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{},
			IPAddress: &armcontainerinstance.IPAddress{
				Type: to.Ptr(armcontainerinstance.ContainerGroupIPAddressTypePublic),
				Ports: []*armcontainerinstance.Port{
					{
						Port:     to.Ptr[int32](80),
						Protocol: to.Ptr(armcontainerinstance.ContainerGroupNetworkProtocolTCP),
					}},
			},
			OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
			Volumes: []*armcontainerinstance.Volume{
				{
					Name: to.Ptr("volume1"),
					AzureFile: &armcontainerinstance.AzureFileVolume{
						ShareName:          to.Ptr("shareName"),
						StorageAccountKey:  to.Ptr("accountKey"),
						StorageAccountName: to.Ptr("accountName"),
					},
				},
				{
					Name:     to.Ptr("volume2"),
					EmptyDir: map[string]any{},
				},
				{
					Name: to.Ptr("volume3"),
					Secret: map[string]*string{
						"secretKey1": to.Ptr("SecretValue1InBase64"),
						"secretKey2": to.Ptr("SecretValue2InBase64"),
					},
				}},
		},
		Location: to.Ptr("west us"),
		Zones: []*string{
			to.Ptr("1")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ContainerGroupProfile = armcontainerinstance.ContainerGroupProfile{
	// 	Properties: &armcontainerinstance.ContainerGroupProfilePropertiesProperties{
	// 		Containers: []*armcontainerinstance.Container{
	// 			{
	// 				Name: to.Ptr("demo1"),
	// 				Properties: &armcontainerinstance.ContainerProperties{
	// 					Command: []*string{
	// 					},
	// 					EnvironmentVariables: []*armcontainerinstance.EnvironmentVariable{
	// 					},
	// 					Image: to.Ptr("nginx"),
	// 					Ports: []*armcontainerinstance.ContainerPort{
	// 						{
	// 							Port: to.Ptr[int32](80),
	// 					}},
	// 					Resources: &armcontainerinstance.ResourceRequirements{
	// 						Requests: &armcontainerinstance.ResourceRequests{
	// 							CPU: to.Ptr[float64](1),
	// 							Gpu: &armcontainerinstance.GpuResource{
	// 								Count: to.Ptr[int32](1),
	// 								SKU: to.Ptr(armcontainerinstance.GpuSKUK80),
	// 							},
	// 							MemoryInGB: to.Ptr[float64](1.5),
	// 						},
	// 					},
	// 					VolumeMounts: []*armcontainerinstance.VolumeMount{
	// 						{
	// 							Name: to.Ptr("volume1"),
	// 							MountPath: to.Ptr("/mnt/volume1"),
	// 							ReadOnly: to.Ptr(false),
	// 						},
	// 						{
	// 							Name: to.Ptr("volume2"),
	// 							MountPath: to.Ptr("/mnt/volume2"),
	// 							ReadOnly: to.Ptr(false),
	// 						},
	// 						{
	// 							Name: to.Ptr("volume3"),
	// 							MountPath: to.Ptr("/mnt/volume3"),
	// 							ReadOnly: to.Ptr(true),
	// 					}},
	// 				},
	// 		}},
	// 		Diagnostics: &armcontainerinstance.ContainerGroupDiagnostics{
	// 			LogAnalytics: &armcontainerinstance.LogAnalytics{
	// 				LogType: to.Ptr(armcontainerinstance.LogAnalyticsLogTypeContainerInsights),
	// 				Metadata: map[string]*string{
	// 					"pod-uuid": to.Ptr("test-metadata-value"),
	// 				},
	// 				WorkspaceID: to.Ptr("workspaceid"),
	// 			},
	// 		},
	// 		ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{
	// 		},
	// 		IPAddress: &armcontainerinstance.IPAddress{
	// 			Type: to.Ptr(armcontainerinstance.ContainerGroupIPAddressTypePublic),
	// 			IP: to.Ptr("10.0.0.1"),
	// 			Ports: []*armcontainerinstance.Port{
	// 				{
	// 					Port: to.Ptr[int32](80),
	// 					Protocol: to.Ptr(armcontainerinstance.ContainerGroupNetworkProtocolTCP),
	// 			}},
	// 		},
	// 		OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
	// 		Revision: to.Ptr[int32](0),
	// 		SKU: to.Ptr(armcontainerinstance.ContainerGroupSKUStandard),
	// 		Volumes: []*armcontainerinstance.Volume{
	// 			{
	// 				Name: to.Ptr("volume1"),
	// 				AzureFile: &armcontainerinstance.AzureFileVolume{
	// 					ShareName: to.Ptr("shareName"),
	// 					StorageAccountName: to.Ptr("accountName"),
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("volume2"),
	// 				EmptyDir: map[string]any{
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("volume3"),
	// 				Secret: map[string]*string{
	// 				},
	// 		}},
	// 	},
	// 	Name: to.Ptr("demo1"),
	// 	Type: to.Ptr("Microsoft.ContainerInstance/containerGroupProfiles"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/demo/providers/Microsoft.ContainerInstance/containerGroupProfiles/demo1"),
	// 	Location: to.Ptr("WestUs"),
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}
