const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Log Analytics clusters in a subscription.
 *
 * @summary Gets the Log Analytics clusters in a subscription.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersSubscriptionList.json
 */
async function clustersSubscriptionList() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.clusters.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

clustersSubscriptionList().catch(console.error);
