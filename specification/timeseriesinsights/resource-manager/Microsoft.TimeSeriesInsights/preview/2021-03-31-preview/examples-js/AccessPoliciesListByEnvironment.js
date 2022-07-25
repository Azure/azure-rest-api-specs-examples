const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available access policies associated with the environment.
 *
 * @summary Lists all the available access policies associated with the environment.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/AccessPoliciesListByEnvironment.json
 */
async function accessPoliciesByEnvironment() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const environmentName = "env1";
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.accessPolicies.listByEnvironment(resourceGroupName, environmentName);
  console.log(result);
}

accessPoliciesByEnvironment().catch(console.error);
