const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves a list of all the built-in policy definition versions for the given policy definition.
 *
 * @summary This operation retrieves a list of all the built-in policy definition versions for the given policy definition.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/listBuiltInPolicyDefinitionVersions.json
 */
async function listBuiltInPolicyDefinitionVersions() {
  const policyDefinitionName = "06a78e20-9358-41c9-923c-fb736d382a12";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const resArray = new Array();
  for await (const item of client.policyDefinitionVersions.listBuiltIn(policyDefinitionName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
