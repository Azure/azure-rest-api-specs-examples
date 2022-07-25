const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an access policy in the specified environment.
 *
 * @summary Create or update an access policy in the specified environment.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/AccessPoliciesCreate.json
 */
async function accessPoliciesCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const environmentName = "env1";
  const accessPolicyName = "ap1";
  const parameters = {
    description: "some description",
    principalObjectId: "aGuid",
    roles: ["Reader"],
  };
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.accessPolicies.createOrUpdate(
    resourceGroupName,
    environmentName,
    accessPolicyName,
    parameters
  );
  console.log(result);
}

accessPoliciesCreate().catch(console.error);
