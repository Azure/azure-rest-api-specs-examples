const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to  This operation creates or updates a variable with the given subscription and name. Policy variables can only be used by a policy definition at the scope they are created or below.
 *
 * @summary  This operation creates or updates a variable with the given subscription and name. Policy variables can only be used by a policy definition at the scope they are created or below.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/createOrUpdateVariable.json
 */
async function createOrUpdateAVariable() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const variableName = "DemoTestVariable";
  const parameters = { columns: [{ columnName: "TestColumn" }] };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.variables.createOrUpdate(variableName, parameters);
  console.log(result);
}
