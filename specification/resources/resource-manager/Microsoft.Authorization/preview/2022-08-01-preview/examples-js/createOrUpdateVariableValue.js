const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to  This operation creates or updates a variable value with the given subscription and name for a given variable. Variable values are scoped to the variable for which they are created for.
 *
 * @summary  This operation creates or updates a variable value with the given subscription and name for a given variable. Variable values are scoped to the variable for which they are created for.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/createOrUpdateVariableValue.json
 */
async function createOrUpdateAVariableValue() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const variableName = "DemoTestVariable";
  const variableValueName = "TestValue";
  const parameters = {
    values: [
      { columnName: "StringColumn", columnValue: "SampleValue" },
      { columnName: "IntegerColumn", columnValue: 10 },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.variableValues.createOrUpdate(
    variableName,
    variableValueName,
    parameters,
  );
  console.log(result);
}
