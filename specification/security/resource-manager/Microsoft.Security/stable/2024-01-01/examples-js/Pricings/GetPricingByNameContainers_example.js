const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the Defender plans pricing configurations of the selected scope (valid scopes are resource id or a subscription id). At the resource level, supported resource types are 'VirtualMachines, VMSS and ARC Machines'.
 *
 * @summary Get the Defender plans pricing configurations of the selected scope (valid scopes are resource id or a subscription id). At the resource level, supported resource types are 'VirtualMachines, VMSS and ARC Machines'.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2024-01-01/examples/Pricings/GetPricingByNameContainers_example.json
 */
async function getPricingsOnSubscriptionContainersPlan() {
  const scopeId = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const pricingName = "Containers";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.pricings.get(scopeId, pricingName);
  console.log(result);
}
