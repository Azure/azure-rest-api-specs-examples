const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a Stream Analytics Cluster or replaces an already existing cluster.
 *
 * @summary Creates a Stream Analytics Cluster or replaces an already existing cluster.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Cluster_Create.json
 */
async function createANewCluster() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "sjrg";
  const clusterName = "An Example Cluster";
  const cluster = {
    location: "North US",
    sku: { name: "Default", capacity: 48 },
    tags: { key: "value" },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    cluster
  );
  console.log(result);
}

createANewCluster().catch(console.error);
