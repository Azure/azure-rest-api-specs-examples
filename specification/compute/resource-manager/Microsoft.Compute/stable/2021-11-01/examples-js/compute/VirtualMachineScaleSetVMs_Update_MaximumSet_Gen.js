const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineScaleSetVMSUpdateMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaa";
  const instanceId = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const parameters = {
    additionalCapabilities: { hibernationEnabled: true, ultraSSDEnabled: true },
    availabilitySet: {
      id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
    },
    diagnosticsProfile: {
      bootDiagnostics: { enabled: true, storageUri: "aaaaaaaaaaaaa" },
    },
    hardwareProfile: {
      vmSize: "Basic_A0",
      vmSizeProperties: { vCPUsAvailable: 9, vCPUsPerCore: 12 },
    },
    instanceView: {
      bootDiagnostics: {
        status: {
          code: "aaaaaaaaaaaaaaaaaaaaaaa",
          displayStatus: "aaaaaa",
          level: "Info",
          message: "a",
          time: new Date("2021-11-30T12:58:26.522Z"),
        },
      },
      disks: [
        {
          name: "aaaaaaaaaaa",
          encryptionSettings: [
            {
              diskEncryptionKey: {
                secretUrl: "aaaaaaaa",
                sourceVault: {
                  id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
                },
              },
              enabled: true,
              keyEncryptionKey: {
                keyUrl: "aaaaaaaaaaaaaa",
                sourceVault: {
                  id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
                },
              },
            },
          ],
          statuses: [
            {
              code: "aaaaaaaaaaaaaaaaaaaaaaa",
              displayStatus: "aaaaaa",
              level: "Info",
              message: "a",
              time: new Date("2021-11-30T12:58:26.522Z"),
            },
          ],
        },
      ],
      maintenanceRedeployStatus: {
        isCustomerInitiatedMaintenanceAllowed: true,
        lastOperationMessage: "aaaaaa",
        lastOperationResultCode: "None",
        maintenanceWindowEndTime: new Date("2021-11-30T12:58:26.531Z"),
        maintenanceWindowStartTime: new Date("2021-11-30T12:58:26.531Z"),
        preMaintenanceWindowEndTime: new Date("2021-11-30T12:58:26.531Z"),
        preMaintenanceWindowStartTime: new Date("2021-11-30T12:58:26.531Z"),
      },
      placementGroupId: "aaa",
      platformFaultDomain: 14,
      platformUpdateDomain: 23,
      rdpThumbPrint: "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
      statuses: [
        {
          code: "aaaaaaaaaaaaaaaaaaaaaaa",
          displayStatus: "aaaaaa",
          level: "Info",
          message: "a",
          time: new Date("2021-11-30T12:58:26.522Z"),
        },
      ],
      vmAgent: {
        extensionHandlers: [
          {
            type: "aaaaaaaaaaaaa",
            status: {
              code: "aaaaaaaaaaaaaaaaaaaaaaa",
              displayStatus: "aaaaaa",
              level: "Info",
              message: "a",
              time: new Date("2021-11-30T12:58:26.522Z"),
            },
            typeHandlerVersion: "aaaaa",
          },
        ],
        statuses: [
          {
            code: "aaaaaaaaaaaaaaaaaaaaaaa",
            displayStatus: "aaaaaa",
            level: "Info",
            message: "a",
            time: new Date("2021-11-30T12:58:26.522Z"),
          },
        ],
        vmAgentVersion: "aaaaaaaaaaaaaaaaaaaaaaa",
      },
      vmHealth: {
        status: {
          code: "aaaaaaaaaaaaaaaaaaaaaaa",
          displayStatus: "aaaaaa",
          level: "Info",
          message: "a",
          time: new Date("2021-11-30T12:58:26.522Z"),
        },
      },
      extensions: [
        {
          name: "aaaaaaaaaaaaaaaaa",
          type: "aaaaaaaaa",
          statuses: [
            {
              code: "aaaaaaaaaaaaaaaaaaaaaaa",
              displayStatus: "aaaaaa",
              level: "Info",
              message: "a",
              time: new Date("2021-11-30T12:58:26.522Z"),
            },
          ],
          substatuses: [
            {
              code: "aaaaaaaaaaaaaaaaaaaaaaa",
              displayStatus: "aaaaaa",
              level: "Info",
              message: "a",
              time: new Date("2021-11-30T12:58:26.522Z"),
            },
          ],
          typeHandlerVersion: "aaaaaaaaaaaaaaaaaaaaaaaaaa",
        },
      ],
    },
    licenseType: "aaaaaaaaaa",
    location: "westus",
    networkProfile: {
      networkApiVersion: "2020-11-01",
      networkInterfaceConfigurations: [
        {
          name: "aaaaaaaaaaa",
          deleteOption: "Delete",
          dnsSettings: { dnsServers: ["aaaaaa"] },
          dscpConfiguration: {
            id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
          },
          enableAcceleratedNetworking: true,
          enableFpga: true,
          enableIPForwarding: true,
          ipConfigurations: [
            {
              name: "aa",
              applicationGatewayBackendAddressPools: [
                {
                  id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
                },
              ],
              applicationSecurityGroups: [
                {
                  id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
                },
              ],
              loadBalancerBackendAddressPools: [
                {
                  id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
                },
              ],
              primary: true,
              privateIPAddressVersion: "IPv4",
              publicIPAddressConfiguration: {
                name: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                deleteOption: "Delete",
                dnsSettings: { domainNameLabel: "aaaaaaaaaaaaaaaaaaaaaaaaa" },
                idleTimeoutInMinutes: 2,
                ipTags: [
                  {
                    ipTagType: "aaaaaaaaaaaaaaaaaaaaaaaaa",
                    tag: "aaaaaaaaaaaaaaaaaaaa",
                  },
                ],
                publicIPAddressVersion: "IPv4",
                publicIPAllocationMethod: "Dynamic",
                publicIPPrefix: {
                  id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
                },
                sku: { name: "Basic", tier: "Regional" },
              },
              subnet: {
                id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
              },
            },
          ],
          networkSecurityGroup: {
            id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
          },
          primary: true,
        },
      ],
      networkInterfaces: [
        {
          deleteOption: "Delete",
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/0/networkInterfaces/vmsstestnetconfig5415",
          primary: true,
        },
      ],
    },
    networkProfileConfiguration: {
      networkInterfaceConfigurations: [
        {
          name: "vmsstestnetconfig5415",
          deleteOption: "Delete",
          dnsSettings: { dnsServers: [] },
          enableAcceleratedNetworking: true,
          enableFpga: true,
          enableIPForwarding: true,
          id: "aaaaaaaa",
          ipConfigurations: [
            {
              name: "vmsstestnetconfig9693",
              applicationGatewayBackendAddressPools: [
                {
                  id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
                },
              ],
              applicationSecurityGroups: [
                {
                  id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
                },
              ],
              id: "aaaaaaaaa",
              loadBalancerBackendAddressPools: [
                {
                  id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
                },
              ],
              loadBalancerInboundNatPools: [
                {
                  id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
                },
              ],
              primary: true,
              privateIPAddressVersion: "IPv4",
              publicIPAddressConfiguration: {
                name: "aaaaaaaaaaaaaaaaaa",
                deleteOption: "Delete",
                dnsSettings: { domainNameLabel: "aaaaaaaaaaaaaaaaaa" },
                idleTimeoutInMinutes: 18,
                ipTags: [{ ipTagType: "aaaaaaa", tag: "aaaaaaaaaaaaaaaaaaaaaaaaaaa" }],
                publicIPAddressVersion: "IPv4",
                publicIPPrefix: {
                  id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
                },
                sku: { name: "Basic", tier: "Regional" },
              },
              subnet: {
                id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/vn4071/subnets/sn5503",
              },
            },
          ],
          networkSecurityGroup: {
            id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
          },
          primary: true,
        },
      ],
    },
    osProfile: {
      adminPassword: "aaaaaaaaaaaaaaaa",
      adminUsername: "Foo12",
      allowExtensionOperations: true,
      computerName: "test000000",
      customData: "aaaa",
      linuxConfiguration: {
        disablePasswordAuthentication: true,
        patchSettings: {
          assessmentMode: "ImageDefault",
          patchMode: "ImageDefault",
        },
        provisionVMAgent: true,
        ssh: { publicKeys: [{ path: "aaa", keyData: "aaaaaa" }] },
      },
      requireGuestProvisionSignal: true,
      secrets: [],
      windowsConfiguration: {
        additionalUnattendContent: [
          {
            componentName: "Microsoft-Windows-Shell-Setup",
            content: "aaaaaaaaaaaaaaaaaaaa",
            passName: "OobeSystem",
            settingName: "AutoLogon",
          },
        ],
        enableAutomaticUpdates: true,
        patchSettings: {
          assessmentMode: "ImageDefault",
          enableHotpatching: true,
          patchMode: "Manual",
        },
        provisionVMAgent: true,
        timeZone: "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
        winRM: {
          listeners: [{ certificateUrl: "aaaaaaaaaaaaaaaaaaaaaa", protocol: "Http" }],
        },
      },
    },
    plan: {
      name: "aaaaaaaaaa",
      product: "aaaaaaaaaaaaaaaaaaaa",
      promotionCode: "aaaaaaaaaaaaaaaaaaaa",
      publisher: "aaaaaaaaaaaaaaaaaaaaaa",
    },
    protectionPolicy: {
      protectFromScaleIn: true,
      protectFromScaleSetActions: true,
    },
    securityProfile: {
      encryptionAtHost: true,
      securityType: "TrustedLaunch",
      uefiSettings: { secureBootEnabled: true, vTpmEnabled: true },
    },
    sku: { name: "Classic", capacity: 29, tier: "aaaaaaaaaaaaaa" },
    storageProfile: {
      dataDisks: [
        {
          name: "vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d",
          caching: "None",
          createOption: "Empty",
          deleteOption: "Delete",
          detachOption: "ForceDetach",
          diskSizeGB: 128,
          image: {
            uri: "https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd",
          },
          lun: 1,
          managedDisk: {
            diskEncryptionSet: { id: "aaaaaaaaaaaa" },
            id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d",
            storageAccountType: "Standard_LRS",
          },
          toBeDetached: true,
          vhd: {
            uri: "https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd",
          },
          writeAcceleratorEnabled: true,
        },
      ],
      imageReference: {
        id: "a",
        offer: "WindowsServer",
        publisher: "MicrosoftWindowsServer",
        sharedGalleryImageId: "aaaaaaaaaaaaaaaaaaaa",
        sku: "2012-R2-Datacenter",
        version: "4.127.20180315",
      },
      osDisk: {
        name: "vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc",
        caching: "None",
        createOption: "FromImage",
        deleteOption: "Delete",
        diffDiskSettings: { option: "Local", placement: "CacheDisk" },
        diskSizeGB: 127,
        encryptionSettings: {
          diskEncryptionKey: {
            secretUrl: "aaaaaaaa",
            sourceVault: {
              id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
            },
          },
          enabled: true,
          keyEncryptionKey: {
            keyUrl: "aaaaaaaaaaaaaa",
            sourceVault: {
              id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
            },
          },
        },
        image: {
          uri: "https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd",
        },
        managedDisk: {
          diskEncryptionSet: { id: "aaaaaaaaaaaa" },
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc",
          storageAccountType: "Standard_LRS",
        },
        osType: "Windows",
        vhd: {
          uri: "https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd",
        },
        writeAcceleratorEnabled: true,
      },
    },
    tags: {},
    userData: "RXhhbXBsZSBVc2VyRGF0YQ==",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.beginUpdateAndWait(
    resourceGroupName,
    vmScaleSetName,
    instanceId,
    parameters
  );
  console.log(result);
}

virtualMachineScaleSetVMSUpdateMaximumSetGen().catch(console.error);
