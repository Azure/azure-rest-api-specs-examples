const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Create_LinuxVmWithAutomaticByPlatformSettings.json
 */
async function createALinuxVmWithAPatchSettingPatchModeOfAutomaticByPlatformAndAutomaticByPlatformSettings() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "myVM", {
    location: "westus",
    hardwareProfile: { vmSize: "Standard_D2s_v3" },
    storageProfile: {
      imageReference: {
        sku: "16.04-LTS",
        publisher: "Canonical",
        version: "latest",
        offer: "UbuntuServer",
      },
      osDisk: {
        caching: "ReadWrite",
        managedDisk: { storageAccountType: "Premium_LRS" },
        name: "myVMosdisk",
        createOption: "FromImage",
      },
    },
    osProfile: {
      adminUsername: "{your-username}",
      computerName: "myVM",
      adminPassword: "{your-password}",
      linuxConfiguration: {
        provisionVMAgent: true,
        patchSettings: {
          patchMode: "AutomaticByPlatform",
          assessmentMode: "AutomaticByPlatform",
          automaticByPlatformSettings: {
            rebootSetting: "Never",
            bypassPlatformSafetyChecksOnUserSchedule: true,
          },
        },
      },
    },
    networkProfile: {
      networkInterfaces: [
        {
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
          primary: true,
        },
      ],
    },
  });
  console.log(result);
}
