const { MicrosoftElastic } = require("@azure/arm-elastic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete traffic filter from the account.
 *
 * @summary Delete traffic filter from the account.
 * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/TrafficFilters_Delete.json
 */
async function trafficFiltersDelete() {
  const subscriptionId =
    process.env["ELASTIC_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ELASTIC_RESOURCE_GROUP"] || "myResourceGroup";
  const monitorName = "myMonitor";
  const rulesetId = "31d91b5afb6f4c2eaaf104c97b1991dd";
  const options = { rulesetId };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftElastic(credential, subscriptionId);
  const result = await client.trafficFilters.delete(resourceGroupName, monitorName, options);
  console.log(result);
}
