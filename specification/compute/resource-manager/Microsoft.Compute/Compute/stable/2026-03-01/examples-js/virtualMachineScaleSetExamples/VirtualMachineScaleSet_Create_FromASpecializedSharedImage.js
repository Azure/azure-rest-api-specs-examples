const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a VM scale set.
 *
 * @summary create or update a VM scale set.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_FromASpecializedSharedImage.json
 */
async function createAScaleSetFromASpecializedSharedImage() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.createOrUpdate(
    "myResourceGroup",
    "{vmss-name}",
    {
      sku: { tier: "Standard", capacity: 3, name: "Standard_D1_v2" },
      location: "westus",
      overprovision: true,
      virtualMachineProfile: {
        storageProfile: {
          imageReference: {
            id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/mySharedGallery/images/mySharedImage",
          },
          osDisk: {
            caching: "ReadWrite",
            managedDisk: { storageAccountType: "Standard_LRS" },
            createOption: "FromImage",
          },
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
      upgradePolicy: { mode: "Manual" },
    },
  );
  console.log(result);
}
