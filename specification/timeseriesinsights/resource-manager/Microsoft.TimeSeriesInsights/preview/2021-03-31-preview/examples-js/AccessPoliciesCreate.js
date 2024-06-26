const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Create or update an access policy in the specified environment.
 *
 * @summary Create or update an access policy in the specified environment.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/AccessPoliciesCreate.json
 */
async function accessPoliciesCreate() {
  const subscriptionId = process.env["TIMESERIESINSIGHTS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["TIMESERIESINSIGHTS_RESOURCE_GROUP"] || "rg1";
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

async function main() {
  accessPoliciesCreate();
}

main().catch(console.error);
