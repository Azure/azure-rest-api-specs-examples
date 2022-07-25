const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available environments associated with the subscription and within the specified resource group.
 *
 * @summary Lists all the available environments associated with the subscription and within the specified resource group.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/EnvironmentsListByResourceGroup.json
 */
async function environmentsByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.environments.listByResourceGroup(resourceGroupName);
  console.log(result);
}

environmentsByResourceGroup().catch(console.error);
