const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateAWindowsVmWithPatchSettingModeOfAutomaticByPlatformAndEnableHotPatchingTrue.json
 */
async function createAWindowsVMWithAPatchSettingPatchModeOfAutomaticByPlatformAndEnableHotpatchingSetToTrue() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmName = "myVM";
  const parameters = {
    hardwareProfile: { vmSize: "Standard_D1_v2" },
    location: "westus",
    networkProfile: {
      networkInterfaces: [
        {
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
          primary: true,
        },
      ],
    },
    osProfile: {
      adminPassword: "{your-password}",
      adminUsername: "{your-username}",
      computerName: "myVM",
      windowsConfiguration: {
        enableAutomaticUpdates: true,
        patchSettings: {
          enableHotpatching: true,
          patchMode: "AutomaticByPlatform",
        },
        provisionVMAgent: true,
      },
    },
    storageProfile: {
      imageReference: {
        offer: "WindowsServer",
        publisher: "MicrosoftWindowsServer",
        sku: "2016-Datacenter",
        version: "latest",
      },
      osDisk: {
        name: "myVMosdisk",
        caching: "ReadWrite",
        createOption: "FromImage",
        managedDisk: { storageAccountType: "Premium_LRS" },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vmName,
    parameters
  );
  console.log(result);
}

createAWindowsVMWithAPatchSettingPatchModeOfAutomaticByPlatformAndEnableHotpatchingSetToTrue().catch(
  console.error
);
