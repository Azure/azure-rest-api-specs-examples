const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the event source with the specified name in the specified subscription, resource group, and environment
 *
 * @summary Deletes the event source with the specified name in the specified subscription, resource group, and environment
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/EventSourcesDelete.json
 */
async function deleteEventSource() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const environmentName = "env1";
  const eventSourceName = "es1";
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.eventSources.delete(
    resourceGroupName,
    environmentName,
    eventSourceName
  );
  console.log(result);
}

deleteEventSource().catch(console.error);
