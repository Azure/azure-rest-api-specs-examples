const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a VM scale set.
 *
 * @summary create or update a VM scale set.
 * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithProxyAgentSettings.json
 */
async function createAScaleSetWithProxyAgentSettingsOfEnabledAndMode() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.createOrUpdate(
    "myResourceGroup",
    "{vmss-name}",
    {
      sku: { tier: "Standard", capacity: 3, name: "Standard_D2s_v3" },
      overprovision: true,
      virtualMachineProfile: {
        storageProfile: {
          imageReference: {
            publisher: "Canonical",
            offer: "0001-com-ubuntu-server-jammy",
            sku: "22_04-lts-gen2",
            version: "latest",
          },
          osDisk: {
            caching: "ReadOnly",
            managedDisk: { storageAccountType: "StandardSSD_LRS" },
            createOption: "FromImage",
          },
        },
        securityProfile: {
          proxyAgentSettings: {
            enabled: true,
            addProxyAgentExtension: true,
            wireServer: {
              inVMAccessControlProfileReferenceId:
                "/subscriptions/{sub-id}/resourceGroups/{rg}/providers/Microsoft.Compute/galleries/{gallery-name}/inVMAccessControlProfiles/{profile-name}/versions/{version}",
            },
            imds: {
              inVMAccessControlProfileReferenceId:
                "/subscriptions/{sub-id}/resourceGroups/{rg}/providers/Microsoft.Compute/galleries/{gallery-name}/inVMAccessControlProfiles/{profile-name}/versions/{version}",
            },
          },
        },
        osProfile: { computerNamePrefix: "{vmss-name}", adminUsername: "{your-username}" },
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
