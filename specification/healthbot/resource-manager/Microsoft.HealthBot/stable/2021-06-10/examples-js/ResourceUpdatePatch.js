const { HealthbotClient } = require("@azure/arm-healthbot");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch a HealthBot.
 *
 * @summary Patch a HealthBot.
 * x-ms-original-file: specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2021-06-10/examples/ResourceUpdatePatch.json
 */
async function botUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "healthbotClient";
  const botName = "samplebotname";
  const parameters = { sku: { name: "F0" } };
  const credential = new DefaultAzureCredential();
  const client = new HealthbotClient(credential, subscriptionId);
  const result = await client.bots.update(resourceGroupName, botName, parameters);
  console.log(result);
}

botUpdate().catch(console.error);
