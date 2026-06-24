const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a VM scale set.
 *
 * @summary create or update a VM scale set.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithLifecycleHooksProfile.json
 */
async function createAVirtualMachineScaleSetWithLifecycleHooksProfile() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "2e2e3046-f85f-4966-8fd2-5fd7bf6ea717";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.createOrUpdate(
    "myResourceGroup",
    "{vmss-name}",
    {
      sku: { tier: "Standard", capacity: 4, name: "Standard_D1_v2" },
      location: "westus",
      lifecycleHooksProfile: {
        lifecycleHooks: [
          {
            type: "UpgradeAutoOSRollingBatchStarting",
            waitDuration: "PT20M",
            defaultAction: "Approve",
          },
        ],
      },
      virtualMachineProfile: {
        storageProfile: {
          imageReference: {
            sku: "2016-Datacenter",
            publisher: "MicrosoftWindowsServer",
            version: "latest",
            offer: "WindowsServer",
          },
          osDisk: {
            caching: "ReadWrite",
            managedDisk: { storageAccountType: "Standard_LRS" },
            createOption: "FromImage",
          },
        },
        osProfile: {
          computerNamePrefix: "{vmss-name}",
          adminUsername: "{your-username}",
          adminPassword: "{your-password}",
        },
        networkProfile: {
          networkInterfaceConfigurations: [
            {
              name: "{vmss-name}",
              primary: true,
              enableIPForwarding: true,
              ipConfigurations: [
                {
                  name: "{vmss-name}",
                  subnet: {
                    id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}",
                  },
                },
              ],
            },
          ],
        },
      },
    },
  );
  console.log(result);
}
