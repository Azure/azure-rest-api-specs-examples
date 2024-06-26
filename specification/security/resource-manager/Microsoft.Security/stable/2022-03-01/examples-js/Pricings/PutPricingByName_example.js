const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a provided Microsoft Defender for Cloud pricing configuration in the subscription.
 *
 * @summary Updates a provided Microsoft Defender for Cloud pricing configuration in the subscription.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2022-03-01/examples/Pricings/PutPricingByName_example.json
 */
async function updatePricingOnSubscription() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const pricingName = "VirtualMachines";
  const pricing = { pricingTier: "Standard", subPlan: "P2" };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.pricings.update(pricingName, pricing);
  console.log(result);
}
