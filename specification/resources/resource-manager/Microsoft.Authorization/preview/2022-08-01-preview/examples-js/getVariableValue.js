const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves a single variable value; given its name, subscription it was created at and the variable it's created for.
 *
 * @summary This operation retrieves a single variable value; given its name, subscription it was created at and the variable it's created for.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/getVariableValue.json
 */
async function retrieveAVariableValue() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const variableName = "DemoTestVariable";
  const variableValueName = "TestValue";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.variableValues.get(variableName, variableValueName);
  console.log(result);
}
