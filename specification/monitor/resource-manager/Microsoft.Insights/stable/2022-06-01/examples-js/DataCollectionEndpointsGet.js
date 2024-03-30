const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the specified data collection endpoint.
 *
 * @summary Returns the specified data collection endpoint.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/DataCollectionEndpointsGet.json
 */
async function getDataCollectionEndpoint() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "703362b3-f278-4e4b-9179-c76eaf41ffc2";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "myResourceGroup";
  const dataCollectionEndpointName = "myCollectionEndpoint";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.dataCollectionEndpoints.get(
    resourceGroupName,
    dataCollectionEndpointName,
  );
  console.log(result);
}
