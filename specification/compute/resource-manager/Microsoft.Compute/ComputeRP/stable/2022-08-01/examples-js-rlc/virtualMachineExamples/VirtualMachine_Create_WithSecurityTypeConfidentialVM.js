const createComputeManagementClient = require("@azure-rest/arm-compute").default,
  { getLongRunningPoller } = require("@azure-rest/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineExamples/VirtualMachine_Create_WithSecurityTypeConfidentialVM.json
 */
async function createAVMWithSecurityTypeConfidentialVMWithPlatformManagedKeys() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "myResourceGroup";
  const vmName = "myVM";
  const options = {
    body: {
      location: "westus",
      properties: {
        hardwareProfile: { vmSize: "Standard_A5" },
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
        },
        securityProfile: {
          securityType: "ConfidentialVM",
          uefiSettings: { secureBootEnabled: true, vTpmEnabled: true },
        },
        storageProfile: {
          imageReference: {
            offer: "2019-datacenter-cvm",
            publisher: "MicrosoftWindowsServer",
            sku: "windows-cvm",
            version: "17763.2183.2109130127",
          },
          osDisk: {
            name: "myVMosdisk",
            caching: "ReadOnly",
            createOption: "FromImage",
            managedDisk: {
              securityProfile: {
                securityEncryptionType: "DiskWithVMGuestState",
              },
              storageAccountType: "StandardSSD_LRS",
            },
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

createAVMWithSecurityTypeConfidentialVMWithPlatformManagedKeys().catch(console.error);
