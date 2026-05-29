const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get the Defender plans pricing configurations of the selected scope (valid scopes are resource id or a subscription id). At the resource level, supported resource types are 'VirtualMachines, VMSS and ARC Machines'.
 *
 * @summary get the Defender plans pricing configurations of the selected scope (valid scopes are resource id or a subscription id). At the resource level, supported resource types are 'VirtualMachines, VMSS and ARC Machines'.
 * x-ms-original-file: 2024-01-01/Pricings/GetResourcePricingByNameContainers_example.json
 */
async function getPricingsOnResourceContainersPlan() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.pricings.get(
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/demo-containers-rg/providers/Microsoft.ContainerService/managedClusters/demo-aks-cluster",
    "Containers",
  );
  console.log(result);
}
