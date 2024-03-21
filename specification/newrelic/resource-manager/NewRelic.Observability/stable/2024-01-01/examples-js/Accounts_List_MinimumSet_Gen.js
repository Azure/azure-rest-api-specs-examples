const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the existing accounts
 *
 * @summary List all the existing accounts
 * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/Accounts_List_MinimumSet_Gen.json
 */
async function accountsListMinimumSetGen() {
  const subscriptionId =
    process.env["NEWRELICOBSERVABILITY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const userEmail = "ruxvg@xqkmdhrnoo.hlmbpm";
  const location = "egh";
  const credential = new DefaultAzureCredential();
  const client = new NewRelicObservability(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.accounts.list(userEmail, location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
