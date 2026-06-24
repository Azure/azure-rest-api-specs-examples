const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a VM scale set.
 *
 * @summary create or update a VM scale set.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithAutomaticSkuMigrationPolicy.json
 */
async function createAScaleSetWithAutomaticSKUMigrationPolicy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.createOrUpdate(
    "myResourceGroup",
    "{vmss-name}",
    {
      sku: { capacity: 10, name: "Mix" },
      location: "westus",
      singlePlacementGroup: false,
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
        priority: "Spot",
        evictionPolicy: "Deallocate",
        billingProfile: { maxPrice: -1 },
      },
      orchestrationMode: "Flexible",
      priorityMixPolicy: { baseRegularPriorityCount: 4, regularPriorityPercentageAboveBase: 50 },
      skuProfile: {
        vmSizes: [
          { name: "Standard_D8s_v5" },
          { name: "Standard_E16s_v5" },
          { name: "Standard_D2s_v5" },
        ],
        allocationStrategy: "CapacityOptimized",
        automaticSkuMigrationPolicy: { enabled: true },
      },
    },
  );
  console.log(result);
}
