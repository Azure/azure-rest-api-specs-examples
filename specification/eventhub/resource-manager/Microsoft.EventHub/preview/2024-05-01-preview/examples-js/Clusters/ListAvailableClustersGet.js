const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List the quantity of available pre-provisioned Event Hubs Clusters, indexed by Azure region.
 *
 * @summary List the quantity of available pre-provisioned Event Hubs Clusters, indexed by Azure region.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2024-05-01-preview/examples/Clusters/ListAvailableClustersGet.json
 */
async function listAvailableClusters() {
  const subscriptionId =
    process.env["EVENTHUB_SUBSCRIPTION_ID"] || "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.clusters.listAvailableClusterRegion();
  console.log(result);
}
