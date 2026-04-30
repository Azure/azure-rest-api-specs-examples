const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a VM scale set.
 *
 * @summary create or update a VM scale set.
 * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithDiskEncryptionSetResource.json
 */
async function createAScaleSetWithDiskEncryptionSetResourceInOsDiskAndDataDisk() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.createOrUpdate(
    "myResourceGroup",
    "{vmss-name}",
    {
      sku: { tier: "Standard", capacity: 3, name: "Standard_DS1_v2" },
      overprovision: true,
      virtualMachineProfile: {
        storageProfile: {
          imageReference: {
            id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}",
          },
          osDisk: {
            caching: "ReadWrite",
            managedDisk: {
              storageAccountType: "Standard_LRS",
              diskEncryptionSet: {
                id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}",
              },
            },
            createOption: "FromImage",
          },
          dataDisks: [
            {
              caching: "ReadWrite",
              managedDisk: {
                storageAccountType: "Standard_LRS",
                diskEncryptionSet: {
                  id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}",
                },
              },
              diskSizeGB: 1023,
              createOption: "Empty",
              lun: 0,
            },
          ],
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
      upgradePolicy: { mode: "Manual" },
      location: "westus",
    },
  );
  console.log(result);
}
