const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the access policy with the specified name in the specified environment.
 *
 * @summary Gets the access policy with the specified name in the specified environment.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/AccessPoliciesGet.json
 */
async function accessPoliciesGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const environmentName = "env1";
  const accessPolicyName = "ap1";
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.accessPolicies.get(
    resourceGroupName,
    environmentName,
    accessPolicyName
  );
  console.log(result);
}

accessPoliciesGet().catch(console.error);
