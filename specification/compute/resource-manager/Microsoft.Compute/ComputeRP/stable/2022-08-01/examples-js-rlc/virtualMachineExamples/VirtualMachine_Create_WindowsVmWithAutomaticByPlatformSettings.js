const createComputeManagementClient = require("@azure-rest/arm-compute").default,
  { getLongRunningPoller } = require("@azure-rest/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineExamples/VirtualMachine_Create_WindowsVmWithAutomaticByPlatformSettings.json
 */
async function createAWindowsVMWithAPatchSettingPatchModeOfAutomaticByPlatformAndAutomaticByPlatformSettings() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "myResourceGroup";
  const vmName = "myVM";
  const options = {
    body: {
      location: "westus",
      properties: {
        hardwareProfile: { vmSize: "Standard_D1_v2" },
        networkProfile: {
          networkInterfaces: [
            {
              id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
              properties: { primary: true },
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
              assessmentMode: "AutomaticByPlatform",
              automaticByPlatformSettings: { rebootSetting: "Never" },
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
      },
    },
    queryParameters: { "api-version": "2022-08-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}",
      subscriptionId,
      resourceGroupName,
      vmName
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createAWindowsVMWithAPatchSettingPatchModeOfAutomaticByPlatformAndAutomaticByPlatformSettings().catch(
  console.error
);
