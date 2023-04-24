const { MicrosoftElastic } = require("@azure/arm-elastic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Detach traffic filter for the given deployment.
 *
 * @summary Detach traffic filter for the given deployment.
 * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/DetachTrafficFilters_Update.json
 */
async function detachTrafficFilterUpdate() {
  const subscriptionId =
    process.env["ELASTIC_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ELASTIC_RESOURCE_GROUP"] || "myResourceGroup";
  const monitorName = "myMonitor";
  const rulesetId = "31d91b5afb6f4c2eaaf104c97b1991dd";
  const options = { rulesetId };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftElastic(credential, subscriptionId);
  const result = await client.detachTrafficFilter.beginUpdateAndWait(
    resourceGroupName,
    monitorName,
    options
  );
  console.log(result);
}
