const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves the built-in policy set definition with the given name.
 *
 * @summary This operation retrieves the built-in policy set definition with the given name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2024-05-01/examples/getBuiltInPolicySetDefinition.json
 */
async function retrieveABuiltInPolicySetDefinition() {
  const policySetDefinitionName = "1f3afdf9-d0c9-4c3d-847f-89da613e70a8";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policySetDefinitions.getBuiltIn(policySetDefinitionName);
  console.log(result);
}
