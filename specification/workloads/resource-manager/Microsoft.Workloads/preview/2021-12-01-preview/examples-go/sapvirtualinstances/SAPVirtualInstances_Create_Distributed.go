package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Create_Distributed.json
func ExampleSAPVirtualInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armworkloads.NewSAPVirtualInstancesClient("8e17e36c-42e9-4cd5-a078-7b44883414e0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"test-rg",
		"X00",
		armworkloads.SAPVirtualInstance{
			Location: to.Ptr("westcentralus"),
			Tags:     map[string]*string{},
			Properties: &armworkloads.SAPVirtualInstanceProperties{
				Configuration: &armworkloads.DeploymentConfiguration{
					ConfigurationType: to.Ptr(armworkloads.SAPConfigurationTypeDeployment),
					AppLocation:       to.Ptr("eastus"),
					InfrastructureConfiguration: &armworkloads.ThreeTierConfiguration{
						AppResourceGroup: to.Ptr("X00-RG"),
						DeploymentType:   to.Ptr(armworkloads.SAPDeploymentTypeThreeTier),
						ApplicationServer: &armworkloads.ApplicationServerConfiguration{
							InstanceCount: to.Ptr[int64](6),
							SubnetID:      to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/dindurkhya-e2etesting/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
							VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
								ImageReference: &armworkloads.ImageReference{
									Offer:     to.Ptr("RHEL-SAP"),
									Publisher: to.Ptr("RedHat"),
									SKU:       to.Ptr("7.4"),
									Version:   to.Ptr("7.4.2019062505"),
								},
								OSProfile: &armworkloads.OSProfile{
									AdminUsername: to.Ptr("{your-username}"),
									OSConfiguration: &armworkloads.LinuxConfiguration{
										OSType:                        to.Ptr(armworkloads.OSTypeLinux),
										DisablePasswordAuthentication: to.Ptr(true),
										SSH: &armworkloads.SSHConfiguration{
											PublicKeys: []*armworkloads.SSHPublicKey{
												{
													KeyData: to.Ptr("ssh-rsa public key"),
												}},
										},
									},
								},
								VMSize: to.Ptr("Standard_E32ds_v4"),
							},
						},
						CentralServer: &armworkloads.CentralServerConfiguration{
							InstanceCount: to.Ptr[int64](1),
							SubnetID:      to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/dindurkhya-e2etesting/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
							VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
								ImageReference: &armworkloads.ImageReference{
									Offer:     to.Ptr("RHEL-SAP"),
									Publisher: to.Ptr("RedHat"),
									SKU:       to.Ptr("7.4"),
									Version:   to.Ptr("7.4.2019062505"),
								},
								OSProfile: &armworkloads.OSProfile{
									AdminUsername: to.Ptr("{your-username}"),
									OSConfiguration: &armworkloads.LinuxConfiguration{
										OSType:                        to.Ptr(armworkloads.OSTypeLinux),
										DisablePasswordAuthentication: to.Ptr(true),
										SSH: &armworkloads.SSHConfiguration{
											PublicKeys: []*armworkloads.SSHPublicKey{
												{
													KeyData: to.Ptr("ssh-rsa public key"),
												}},
										},
									},
								},
								VMSize: to.Ptr("Standard_E16ds_v4"),
							},
						},
						DatabaseServer: &armworkloads.DatabaseConfiguration{
							DatabaseType:  to.Ptr(armworkloads.SAPDatabaseTypeHANA),
							InstanceCount: to.Ptr[int64](1),
							SubnetID:      to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/dindurkhya-e2etesting/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/dbsubnet"),
							VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
								ImageReference: &armworkloads.ImageReference{
									Offer:     to.Ptr("RHEL-SAP"),
									Publisher: to.Ptr("RedHat"),
									SKU:       to.Ptr("7.4"),
									Version:   to.Ptr("7.4.2019062505"),
								},
								OSProfile: &armworkloads.OSProfile{
									AdminUsername: to.Ptr("{your-username}"),
									OSConfiguration: &armworkloads.LinuxConfiguration{
										OSType:                        to.Ptr(armworkloads.OSTypeLinux),
										DisablePasswordAuthentication: to.Ptr(true),
										SSH: &armworkloads.SSHConfiguration{
											PublicKeys: []*armworkloads.SSHPublicKey{
												{
													KeyData: to.Ptr("ssh-rsa public key"),
												}},
										},
									},
								},
								VMSize: to.Ptr("Standard_M32ts"),
							},
						},
					},
				},
				Environment: to.Ptr(armworkloads.SAPEnvironmentTypeProd),
				SapProduct:  to.Ptr(armworkloads.SAPProductTypeS4HANA),
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
