const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available environments within a subscription, irrespective of the resource groups.
 *
 * @summary Lists all the available environments within a subscription, irrespective of the resource groups.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/EnvironmentsListBySubscription.json
 */
async function environmentsBySubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.environments.listBySubscription();
  console.log(result);
}

environmentsBySubscription().catch(console.error);
