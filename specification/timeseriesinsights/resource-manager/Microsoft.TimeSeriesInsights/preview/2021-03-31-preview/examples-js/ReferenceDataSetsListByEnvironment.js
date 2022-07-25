const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available reference data sets associated with the subscription and within the specified resource group and environment.
 *
 * @summary Lists all the available reference data sets associated with the subscription and within the specified resource group and environment.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/ReferenceDataSetsListByEnvironment.json
 */
async function referenceDataSetsListByEnvironment() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const environmentName = "env1";
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.referenceDataSets.listByEnvironment(
    resourceGroupName,
    environmentName
  );
  console.log(result);
}

referenceDataSetsListByEnvironment().catch(console.error);
