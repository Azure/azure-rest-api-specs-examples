const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the event source with the specified name in the specified subscription, resource group, and environment.
 *
 * @summary Updates the event source with the specified name in the specified subscription, resource group, and environment.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/EventSourcesPatchTags.json
 */
async function updateEventSource() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const environmentName = "env1";
  const eventSourceName = "es1";
  const eventSourceUpdateParameters = {
    tags: { someKey: "someValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.eventSources.update(
    resourceGroupName,
    environmentName,
    eventSourceName,
    eventSourceUpdateParameters
  );
  console.log(result);
}

updateEventSource().catch(console.error);
