const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Updates the environment with the specified name in the specified subscription and resource group.
 *
 * @summary Updates the environment with the specified name in the specified subscription and resource group.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/EnvironmentsPatchTags.json
 */
async function environmentsUpdate() {
  const subscriptionId = process.env["TIMESERIESINSIGHTS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["TIMESERIESINSIGHTS_RESOURCE_GROUP"] || "rg1";
  const environmentName = "env1";
  const environmentUpdateParameters = {
    kind: "Gen1",
    tags: { someTag: "someTagValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.environments.beginUpdateAndWait(
    resourceGroupName,
    environmentName,
    environmentUpdateParameters
  );
  console.log(result);
}

async function main() {
  environmentsUpdate();
}

main().catch(console.error);
