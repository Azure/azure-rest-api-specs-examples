const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateAVmWithUefiSettings.json
 */
async function createAVMWithUefiSettingsOfSecureBootAndVTpm() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmName = "myVM";
  const parameters = {
    hardwareProfile: { vmSize: "Standard_D2s_v3" },
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
    },
    securityProfile: {
      securityType: "TrustedLaunch",
      uefiSettings: { secureBootEnabled: true, vTpmEnabled: true },
    },
    storageProfile: {
      imageReference: {
        offer: "windowsserver-gen2preview-preview",
        publisher: "MicrosoftWindowsServer",
        sku: "windows10-tvm",
        version: "18363.592.2001092016",
      },
      osDisk: {
        name: "myVMosdisk",
        caching: "ReadOnly",
        createOption: "FromImage",
        managedDisk: { storageAccountType: "StandardSSD_LRS" },
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

createAVMWithUefiSettingsOfSecureBootAndVTpm().catch(console.error);
