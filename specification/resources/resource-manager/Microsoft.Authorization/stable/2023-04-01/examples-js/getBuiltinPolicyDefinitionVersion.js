const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves the built-in policy definition version with the given name.
 *
 * @summary This operation retrieves the built-in policy definition version with the given name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/getBuiltinPolicyDefinitionVersion.json
 */
async function retrieveABuiltInPolicyDefinitionVersion() {
  const policyDefinitionName = "7433c107-6db4-4ad1-b57a-a76dce0154a1";
  const policyDefinitionVersion = "1.2.1";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyDefinitionVersions.getBuiltIn(
    policyDefinitionName,
    policyDefinitionVersion,
  );
  console.log(result);
}
