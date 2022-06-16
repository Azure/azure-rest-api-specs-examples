package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/VendorSkuCreate.json
func ExampleVendorSKUsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhybridnetwork.NewVendorSKUsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"TestVendor",
		"TestSku",
		armhybridnetwork.VendorSKU{
			Properties: &armhybridnetwork.VendorSKUPropertiesFormat{
				DeploymentMode:             to.Ptr(armhybridnetwork.SKUDeploymentModePrivateEdgeZone),
				ManagedApplicationTemplate: map[string]interface{}{},
				NetworkFunctionTemplate: &armhybridnetwork.NetworkFunctionTemplate{
					NetworkFunctionRoleConfigurations: []*armhybridnetwork.NetworkFunctionRoleConfiguration{
						{
							CustomProfile: &armhybridnetwork.CustomProfile{
								MetadataConfigurationPath: to.Ptr("/var/logs/network.cfg"),
							},
							NetworkInterfaces: []*armhybridnetwork.NetworkInterface{
								{
									IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
										{
											Gateway:            to.Ptr(""),
											IPAddress:          to.Ptr(""),
											IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
											IPVersion:          to.Ptr(armhybridnetwork.IPVersionIPv4),
											Subnet:             to.Ptr(""),
										}},
									MacAddress:           to.Ptr(""),
									NetworkInterfaceName: to.Ptr("nic1"),
									VMSwitchType:         to.Ptr(armhybridnetwork.VMSwitchTypeWan),
								},
								{
									IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
										{
											Gateway:            to.Ptr(""),
											IPAddress:          to.Ptr(""),
											IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
											IPVersion:          to.Ptr(armhybridnetwork.IPVersionIPv4),
											Subnet:             to.Ptr(""),
										}},
									MacAddress:           to.Ptr(""),
									NetworkInterfaceName: to.Ptr("nic2"),
									VMSwitchType:         to.Ptr(armhybridnetwork.VMSwitchTypeManagement),
								}},
							OSProfile: &armhybridnetwork.OsProfile{
								AdminUsername: to.Ptr("dummyuser"),
								CustomData:    to.Ptr("base-64 encoded string of custom data"),
								LinuxConfiguration: &armhybridnetwork.LinuxConfiguration{
									SSH: &armhybridnetwork.SSHConfiguration{
										PublicKeys: []*armhybridnetwork.SSHPublicKey{
											{
												Path:    to.Ptr("home/user/.ssh/authorized_keys"),
												KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAgEAwrr66r8n6B8Y0zMF3dOpXEapIQD9DiYQ6D6/zwor9o39jSkHNiMMER/GETBbzP83LOcekm02aRjo55ArO7gPPVvCXbrirJu9pkm4AC4BBre5xSLS= user@constoso-DSH"),
											}},
									},
								},
							},
							RoleName: to.Ptr("test"),
							RoleType: to.Ptr(armhybridnetwork.NetworkFunctionRoleConfigurationTypeVirtualMachine),
							StorageProfile: &armhybridnetwork.StorageProfile{
								DataDisks: []*armhybridnetwork.DataDisk{
									{
										Name:         to.Ptr("DataDisk1"),
										CreateOption: to.Ptr(armhybridnetwork.DiskCreateOptionTypesEmpty),
										DiskSizeGB:   to.Ptr[int32](10),
									}},
								ImageReference: &armhybridnetwork.ImageReference{
									Offer:     to.Ptr("UbuntuServer"),
									Publisher: to.Ptr("Canonical"),
									SKU:       to.Ptr("18.04-LTS"),
									Version:   to.Ptr("18.04.201804262"),
								},
								OSDisk: &armhybridnetwork.OsDisk{
									Name:       to.Ptr("vhdName"),
									DiskSizeGB: to.Ptr[int32](30),
									OSType:     to.Ptr(armhybridnetwork.OperatingSystemTypesLinux),
									Vhd: &armhybridnetwork.VirtualHardDisk{
										URI: to.Ptr("https://contoso.net/link/vnd.vhd?sp=rl&st=2020-10-08T20:38:19Z&se=2020-12-09T19:38:00Z&sv=2019-12-12&sr=b&sig=7BM2f4yOw%3D"),
									},
								},
							},
							VirtualMachineSize: to.Ptr(armhybridnetwork.VirtualMachineSizeTypesStandardD3V2),
						}},
				},
				NetworkFunctionType: to.Ptr(armhybridnetwork.NetworkFunctionTypeVirtualNetworkFunction),
				Preview:             to.Ptr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
