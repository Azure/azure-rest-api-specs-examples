const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the available Event Hubs Clusters within an ARM resource group
 *
 * @summary Lists the available Event Hubs Clusters within an ARM resource group
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/Clusters/ClustersListBySubscription.json
 */
async function clustersListBySubscription() {
  const subscriptionId =
    process.env["EVENTHUB_SUBSCRIPTION_ID"] || "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.clusters.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
