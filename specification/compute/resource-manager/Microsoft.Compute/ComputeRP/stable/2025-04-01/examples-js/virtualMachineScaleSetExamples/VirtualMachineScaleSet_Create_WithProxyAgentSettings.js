const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update a VM scale set.
 *
 * @summary Create or update a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithProxyAgentSettings.json
 */
async function createAScaleSetWithProxyAgentSettingsOfEnabledAndMode() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const vmScaleSetName = "{vmss-name}";
  const parameters = {
    location: "westus",
    overprovision: true,
    sku: { name: "Standard_D2s_v3", capacity: 3, tier: "Standard" },
    upgradePolicy: { mode: "Manual" },
    virtualMachineProfile: {
      networkProfile: {
        networkInterfaceConfigurations: [
          {
            name: "{vmss-name}",
            enableIPForwarding: true,
            ipConfigurations: [
              {
                name: "{vmss-name}",
                subnet: {
                  id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}",
                },
              },
            ],
            primary: true,
          },
        ],
      },
      osProfile: {
        adminUsername: "{your-username}",
        computerNamePrefix: "{vmss-name}",
      },
      securityProfile: {
        proxyAgentSettings: {
          addProxyAgentExtension: true,
          enabled: true,
          imds: {
            inVMAccessControlProfileReferenceId:
              "/subscriptions/{sub-id}/resourceGroups/{rg}/providers/Microsoft.Compute/galleries/{gallery-name}/inVMAccessControlProfiles/{profile-name}/versions/{version}",
          },
          wireServer: {
            inVMAccessControlProfileReferenceId:
              "/subscriptions/{sub-id}/resourceGroups/{rg}/providers/Microsoft.Compute/galleries/{gallery-name}/inVMAccessControlProfiles/{profile-name}/versions/{version}",
          },
        },
      },
      storageProfile: {
        imageReference: {
          offer: "0001-com-ubuntu-server-jammy",
          publisher: "Canonical",
          sku: "22_04-lts-gen2",
          version: "latest",
        },
        osDisk: {
          caching: "ReadOnly",
          createOption: "FromImage",
          managedDisk: { storageAccountType: "StandardSSD_LRS" },
        },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vmScaleSetName,
    parameters,
  );
  console.log(result);
}
