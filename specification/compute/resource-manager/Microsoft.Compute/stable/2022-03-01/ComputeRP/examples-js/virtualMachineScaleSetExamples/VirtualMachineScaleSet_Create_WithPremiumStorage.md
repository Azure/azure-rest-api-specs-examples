Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a VM scale set.
 *
 * @summary Create or update a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithPremiumStorage.json
 */
async function createAScaleSetWithPremiumStorage() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmScaleSetName = "{vmss-name}";
  const parameters = {
    location: "westus",
    overprovision: true,
    sku: { name: "Standard_D1_v2", capacity: 3, tier: "Standard" },
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
          offer: "WindowsServer",
          publisher: "MicrosoftWindowsServer",
          sku: "2016-Datacenter",
          version: "latest",
        },
        osDisk: {
          caching: "ReadWrite",
          createOption: "FromImage",
          managedDisk: { storageAccountType: "Premium_LRS" },
        },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vmScaleSetName,
    parameters
  );
  console.log(result);
}

createAScaleSetWithPremiumStorage().catch(console.error);
```
