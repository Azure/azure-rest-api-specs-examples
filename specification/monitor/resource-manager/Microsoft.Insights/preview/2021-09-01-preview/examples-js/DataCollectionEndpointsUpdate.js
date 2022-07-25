const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates part of a data collection endpoint.
 *
 * @summary Updates part of a data collection endpoint.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/DataCollectionEndpointsUpdate.json
 */
async function updateDataCollectionEndpoint() {
  const subscriptionId = "703362b3-f278-4e4b-9179-c76eaf41ffc2";
  const resourceGroupName = "myResourceGroup";
  const dataCollectionEndpointName = "myCollectionEndpoint";
  const body = { tags: { tag1: "A", tag2: "B", tag3: "C" } };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.dataCollectionEndpoints.update(
    resourceGroupName,
    dataCollectionEndpointName,
    options
  );
  console.log(result);
}

updateDataCollectionEndpoint().catch(console.error);
