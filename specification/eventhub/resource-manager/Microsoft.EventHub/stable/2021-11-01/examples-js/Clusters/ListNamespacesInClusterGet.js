const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all Event Hubs Namespace IDs in an Event Hubs Dedicated Cluster.
 *
 * @summary List all Event Hubs Namespace IDs in an Event Hubs Dedicated Cluster.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/Clusters/ListNamespacesInClusterGet.json
 */
async function listNamespacesInCluster() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "myResourceGroup";
  const clusterName = "testCluster";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.clusters.listNamespaces(resourceGroupName, clusterName);
  console.log(result);
}

listNamespacesInCluster().catch(console.error);
