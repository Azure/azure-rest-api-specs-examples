const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a subscription-level key used for Real User Metrics collection.
 *
 * @summary Delete a subscription-level key used for Real User Metrics collection.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/TrafficManagerUserMetricsKeys-DELETE.json
 */
async function trafficManagerUserMetricsKeysDelete() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.trafficManagerUserMetricsKeys.delete();
  console.log(result);
}

trafficManagerUserMetricsKeysDelete().catch(console.error);
