const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists associations for the specified data collection endpoint.
 *
 * @summary Lists associations for the specified data collection endpoint.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/DataCollectionRuleAssociationsListByDataCollectionEndpoint.json
 */
async function listAssociationsForSpecifiedDataCollectionEndpoint() {
  const subscriptionId = "703362b3-f278-4e4b-9179-c76eaf41ffc2";
  const resourceGroupName = "myResourceGroup";
  const dataCollectionEndpointName = "myDataCollectionEndpointName";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dataCollectionRuleAssociations.listByDataCollectionEndpoint(
    resourceGroupName,
    dataCollectionEndpointName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAssociationsForSpecifiedDataCollectionEndpoint().catch(console.error);
