const { SubscriptionClient } = require("@azure/arm-resources-subscriptions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Compares a subscriptions logical zone mapping
 *
 * @summary Compares a subscriptions logical zone mapping
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-12-01/examples/PostCheckZonePeers.json
 */
async function getLogicalZoneMapping() {
  const subscriptionId = "8d65815f-a5b6-402f-9298-045155da7d74";
  const parameters = {
    location: "eastus",
    subscriptionIds: ["subscriptions/11111111-1111-1111-1111-111111111111"],
  };
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const result = await client.subscriptions.checkZonePeers(subscriptionId, parameters);
  console.log(result);
}
