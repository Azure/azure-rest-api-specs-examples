const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a VM scale set.
 *
 * @summary create or update a VM scale set.
 * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithSecurityTypeConfidentialVM.json
 */
async function createAScaleSetWithSecurityTypeAsConfidentialVM() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.createOrUpdate(
    "myResourceGroup",
    "{vmss-name}",
    {
      sku: { tier: "Standard", capacity: 3, name: "Standard_DC2as_v5" },
      overprovision: true,
      virtualMachineProfile: {
        storageProfile: {
          imageReference: {
            sku: "windows-cvm",
            publisher: "MicrosoftWindowsServer",
            version: "17763.2183.2109130127",
            offer: "2019-datacenter-cvm",
          },
          osDisk: {
            caching: "ReadOnly",
            managedDisk: {
              storageAccountType: "StandardSSD_LRS",
              securityProfile: { securityEncryptionType: "VMGuestStateOnly" },
            },
            createOption: "FromImage",
          },
        },
        securityProfile: {
          uefiSettings: { secureBootEnabled: true, vTpmEnabled: true },
          securityType: "ConfidentialVM",
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
