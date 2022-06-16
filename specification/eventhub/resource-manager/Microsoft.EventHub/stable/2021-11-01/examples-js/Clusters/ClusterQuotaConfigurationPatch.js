const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Replace all specified Event Hubs Cluster settings with those contained in the request body. Leaves the settings not specified in the request body unmodified.
 *
 * @summary Replace all specified Event Hubs Cluster settings with those contained in the request body. Leaves the settings not specified in the request body unmodified.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/Clusters/ClusterQuotaConfigurationPatch.json
 */
async function clustersQuotasConfigurationPatch() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const clusterName = "testCluster";
  const parameters = {
    settings: {
      eventhubPerNamespaceQuota: "20",
      namespacesPerClusterQuota: "200",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.configuration.patch(resourceGroupName, clusterName, parameters);
  console.log(result);
}

clustersQuotasConfigurationPatch().catch(console.error);
