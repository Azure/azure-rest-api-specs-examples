```go
package armcontainerinstance_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance"
)

// x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-09-01/examples/ContainerGroupsCreateOrUpdate.json
func ExampleContainerGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerinstance.NewContainerGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<container-group-name>",
		armcontainerinstance.ContainerGroup{
			Resource: armcontainerinstance.Resource{
				Location: to.StringPtr("<location>"),
			},
			Identity: &armcontainerinstance.ContainerGroupIdentity{
				Type: armcontainerinstance.ResourceIdentityTypeSystemAssignedUserAssigned.ToPtr(),
				UserAssignedIdentities: map[string]*armcontainerinstance.Components10Wh5UdSchemasContainergroupidentityPropertiesUserassignedidentitiesAdditionalproperties{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-name": {},
				},
			},
			Properties: &armcontainerinstance.ContainerGroupProperties{
				Containers: []*armcontainerinstance.Container{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armcontainerinstance.ContainerProperties{
							Command:              []*string{},
							EnvironmentVariables: []*armcontainerinstance.EnvironmentVariable{},
							Image:                to.StringPtr("<image>"),
							Ports: []*armcontainerinstance.ContainerPort{
								{
									Port: to.Int32Ptr(80),
								}},
							Resources: &armcontainerinstance.ResourceRequirements{
								Requests: &armcontainerinstance.ResourceRequests{
									CPU: to.Float64Ptr(1),
									Gpu: &armcontainerinstance.GpuResource{
										Count: to.Int32Ptr(1),
										SKU:   armcontainerinstance.GpuSKUK80.ToPtr(),
									},
									MemoryInGB: to.Float64Ptr(1.5),
								},
							},
							VolumeMounts: []*armcontainerinstance.VolumeMount{
								{
									Name:      to.StringPtr("<name>"),
									MountPath: to.StringPtr("<mount-path>"),
									ReadOnly:  to.BoolPtr(false),
								},
								{
									Name:      to.StringPtr("<name>"),
									MountPath: to.StringPtr("<mount-path>"),
									ReadOnly:  to.BoolPtr(false),
								},
								{
									Name:      to.StringPtr("<name>"),
									MountPath: to.StringPtr("<mount-path>"),
									ReadOnly:  to.BoolPtr(true),
								}},
						},
					}},
				Diagnostics: &armcontainerinstance.ContainerGroupDiagnostics{
					LogAnalytics: &armcontainerinstance.LogAnalytics{
						LogType: armcontainerinstance.LogAnalyticsLogTypeContainerInsights.ToPtr(),
						Metadata: map[string]*string{
							"test-key": to.StringPtr("test-metadata-value"),
						},
						WorkspaceID:         to.StringPtr("<workspace-id>"),
						WorkspaceKey:        to.StringPtr("<workspace-key>"),
						WorkspaceResourceID: to.StringPtr("<workspace-resource-id>"),
					},
				},
				DNSConfig: &armcontainerinstance.DNSConfiguration{
					NameServers: []*string{
						to.StringPtr("1.1.1.1")},
					Options:       to.StringPtr("<options>"),
					SearchDomains: to.StringPtr("<search-domains>"),
				},
				ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{},
				IPAddress: &armcontainerinstance.IPAddress{
					Type:         armcontainerinstance.ContainerGroupIPAddressTypePublic.ToPtr(),
					DNSNameLabel: to.StringPtr("<dnsname-label>"),
					Ports: []*armcontainerinstance.Port{
						{
							Port:     to.Int32Ptr(80),
							Protocol: armcontainerinstance.ContainerGroupNetworkProtocolTCP.ToPtr(),
						}},
				},
				OSType: armcontainerinstance.OperatingSystemTypesLinux.ToPtr(),
				SubnetIDs: []*armcontainerinstance.ContainerGroupSubnetID{
					{
						ID: to.StringPtr("<id>"),
					}},
				Volumes: []*armcontainerinstance.Volume{
					{
						Name: to.StringPtr("<name>"),
						AzureFile: &armcontainerinstance.AzureFileVolume{
							ShareName:          to.StringPtr("<share-name>"),
							StorageAccountKey:  to.StringPtr("<storage-account-key>"),
							StorageAccountName: to.StringPtr("<storage-account-name>"),
						},
					},
					{
						Name:     to.StringPtr("<name>"),
						EmptyDir: map[string]interface{}{},
					},
					{
						Name: to.StringPtr("<name>"),
						Secret: map[string]*string{
							"secretKey1": to.StringPtr("SecretValue1InBase64"),
							"secretKey2": to.StringPtr("SecretValue2InBase64"),
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ContainerGroup.ID: %s\n", *res.ID)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerinstance%2Farmcontainerinstance%2Fv0.1.0/sdk/resourcemanager/containerinstance/armcontainerinstance/README.md) on how to add the SDK to your project and authenticate.
