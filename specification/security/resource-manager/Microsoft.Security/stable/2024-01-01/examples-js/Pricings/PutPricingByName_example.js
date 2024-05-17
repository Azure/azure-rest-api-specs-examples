const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a provided Microsoft Defender for Cloud pricing configuration in the scope. Valid scopes are: subscription id or a specific resource id (Supported resources are: 'VirtualMachines, VMSS and ARC Machines' and only for plan='VirtualMachines' and subPlan='P1').
 *
 * @summary Updates a provided Microsoft Defender for Cloud pricing configuration in the scope. Valid scopes are: subscription id or a specific resource id (Supported resources are: 'VirtualMachines, VMSS and ARC Machines' and only for plan='VirtualMachines' and subPlan='P1').
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2024-01-01/examples/Pricings/PutPricingByName_example.json
 */
async function updatePricingOnSubscriptionExampleForCloudPosturePlan() {
  const scopeId = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const pricingName = "CloudPosture";
  const pricing = { pricingTier: "Standard" };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.pricings.update(scopeId, pricingName, pricing);
  console.log(result);
}
