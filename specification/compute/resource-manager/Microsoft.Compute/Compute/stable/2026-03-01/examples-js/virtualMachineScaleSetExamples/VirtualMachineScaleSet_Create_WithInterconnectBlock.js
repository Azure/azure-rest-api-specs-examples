const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a VM scale set.
 *
 * @summary create or update a VM scale set.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithInterconnectBlock.json
 */
async function createOrUpdateAScaleSetWithInterconnectBlock() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.createOrUpdate(
    "myResourceGroup",
    "{vmss-name}",
    {
      sku: { tier: "Standard", capacity: 3, name: "Standard_ND128isr_GB300_v6" },
      location: "westus",
      overprovision: true,
      highSpeedInterconnectPlacement: "Trunk",
      virtualMachineProfile: {
        storageProfile: {
          imageReference: {
            publisher: "microsoft-dsvm",
            offer: "ubuntu-hpc",
            sku: "2404-gb",
            version: "latest",
          },
          osDisk: {
            caching: "ReadWrite",
            managedDisk: { storageAccountType: "Premium_LRS" },
            createOption: "FromImage",
          },
        },
        osProfile: {
          computerNamePrefix: "{vmss-name}",
          adminUsername: "{your-username}",
          adminPassword: "{your-password}",
          linuxConfiguration: { disablePasswordAuthentication: false },
        },
        networkProfile: {
          interconnectGroupProfile: {
            interconnectGroup: {
              id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup",
            },
            subgroups: [
              {
                id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup/subgroups/subgroup0",
              },
            ],
          },
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
        interconnectBlockProfile: {
          interconnectBlock: {
            id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/interconnectBlocks/myInterconnectBlock",
          },
        },
      },
      upgradePolicy: { mode: "Manual" },
      zones: ["1"],
    },
  );
  console.log(result);
}
