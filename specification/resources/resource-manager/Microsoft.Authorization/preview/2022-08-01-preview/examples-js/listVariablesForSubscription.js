const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves the list of all variables associated with the given subscription.
 *
 * @summary This operation retrieves the list of all variables associated with the given subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/listVariablesForSubscription.json
 */
async function listVariablesThatApplyToASubscription() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.variables.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
