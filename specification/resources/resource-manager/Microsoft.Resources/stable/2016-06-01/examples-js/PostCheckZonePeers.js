const { SubscriptionClient } = require("@azure/arm-subscriptions-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Compares a subscriptions logical zone mapping
 *
 * @summary Compares a subscriptions logical zone mapping
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2016-06-01/examples/PostCheckZonePeers.json
 */
async function getLogicalZoneMapping() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000000";
  const parameters = {
    location: "eastus",
    subscriptionIds: ["subscriptions/11111111-1111-1111-1111-111111111111"],
  };
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const result = await client.subscriptions.checkZonePeers(subscriptionId, parameters);
  console.log(result);
}
