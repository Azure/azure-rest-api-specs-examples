const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation deletes a variable, given its name and the subscription it was created in. The scope of a variable is the part of its ID preceding '/providers/Microsoft.Authorization/variables/{variableName}'.
 *
 * @summary This operation deletes a variable, given its name and the subscription it was created in. The scope of a variable is the part of its ID preceding '/providers/Microsoft.Authorization/variables/{variableName}'.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/deleteVariable.json
 */
async function deleteAVariable() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const variableName = "DemoTestVariable";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.variables.delete(variableName);
  console.log(result);
}
