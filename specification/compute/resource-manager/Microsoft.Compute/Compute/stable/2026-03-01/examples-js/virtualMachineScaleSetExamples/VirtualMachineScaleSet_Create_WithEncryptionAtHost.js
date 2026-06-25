const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a VM scale set.
 *
 * @summary create or update a VM scale set.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithEncryptionAtHost.json
 */
async function createAScaleSetWithHostEncryptionUsingEncryptionAtHostProperty() {
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
            sku: "windows2016",
            publisher: "microsoft-ads",
            version: "latest",
            offer: "windows-data-science-vm",
          },
          osDisk: {
            caching: "ReadOnly",
            managedDisk: { storageAccountType: "Standard_LRS" },
            createOption: "FromImage",
          },
        },
        securityProfile: { encryptionAtHost: true },
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
      plan: { publisher: "microsoft-ads", product: "windows-data-science-vm", name: "windows2016" },
      location: "westus",
    },
  );
  console.log(result);
}
