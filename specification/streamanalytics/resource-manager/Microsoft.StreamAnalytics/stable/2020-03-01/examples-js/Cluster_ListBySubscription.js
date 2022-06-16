const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the clusters in the given subscription.
 *
 * @summary Lists all of the clusters in the given subscription.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Cluster_ListBySubscription.json
 */
async function listTheClustersInASubscription() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.clusters.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listTheClustersInASubscription().catch(console.error);
