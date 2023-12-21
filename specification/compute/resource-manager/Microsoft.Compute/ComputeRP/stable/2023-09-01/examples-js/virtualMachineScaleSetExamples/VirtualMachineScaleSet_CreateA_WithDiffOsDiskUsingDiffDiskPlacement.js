const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a VM scale set.
 *
 * @summary Create or update a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_CreateA_WithDiffOsDiskUsingDiffDiskPlacement.json
 */
async function createAScaleSetWithEphemeralOSDisksUsingPlacementProperty() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const vmScaleSetName = "{vmss-name}";
  const parameters = {
    location: "westus",
    overprovision: true,
    plan: {
      name: "windows2016",
      product: "windows-data-science-vm",
      publisher: "microsoft-ads",
    },
    sku: { name: "Standard_DS1_v2", capacity: 3, tier: "Standard" },
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
        adminPassword: "{your-password}",
        adminUsername: "{your-username}",
        computerNamePrefix: "{vmss-name}",
      },
      storageProfile: {
        imageReference: {
          offer: "windows-data-science-vm",
          publisher: "microsoft-ads",
          sku: "windows2016",
          version: "latest",
        },
        osDisk: {
          caching: "ReadOnly",
          createOption: "FromImage",
          diffDiskSettings: { option: "Local", placement: "ResourceDisk" },
          managedDisk: { storageAccountType: "Standard_LRS" },
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
