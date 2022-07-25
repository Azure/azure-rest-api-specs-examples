const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the reference data set with the specified name in the specified subscription, resource group, and environment
 *
 * @summary Deletes the reference data set with the specified name in the specified subscription, resource group, and environment
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/ReferenceDataSetsDelete.json
 */
async function referenceDataSetsDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const environmentName = "env1";
  const referenceDataSetName = "rds1";
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.referenceDataSets.delete(
    resourceGroupName,
    environmentName,
    referenceDataSetName
  );
  console.log(result);
}

referenceDataSetsDelete().catch(console.error);
