const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves the list of all variable values associated with the given variable that is at a subscription level.
 *
 * @summary This operation retrieves the list of all variable values associated with the given variable that is at a subscription level.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/listVariableValuesForSubscription.json
 */
async function listVariableValuesThatApplyToAVariableAtSubscriptionLevel() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const variableName = "DemoTestVariable";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.variableValues.list(variableName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
