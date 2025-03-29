const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation lists all the built-in policy definition versions for all built-in policy definitions.
 *
 * @summary This operation lists all the built-in policy definition versions for all built-in policy definitions.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/listAllBuiltInPolicyDefinitionVersions.json
 */
async function listAllBuiltInPolicyDefinitionVersions() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyDefinitionVersions.listAllBuiltins();
  console.log(result);
}
