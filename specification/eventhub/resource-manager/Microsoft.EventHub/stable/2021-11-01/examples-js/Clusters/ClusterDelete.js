const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing Event Hubs Cluster. This operation is idempotent.
 *
 * @summary Deletes an existing Event Hubs Cluster. This operation is idempotent.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/Clusters/ClusterDelete.json
 */
async function clusterDelete() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "myResourceGroup";
  const clusterName = "testCluster";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginDeleteAndWait(resourceGroupName, clusterName);
  console.log(result);
}

clusterDelete().catch(console.error);
