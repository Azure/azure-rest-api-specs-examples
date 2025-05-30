const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves the built-in policy definition with the given name.
 *
 * @summary This operation retrieves the built-in policy definition with the given name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getBuiltinPolicyDefinition.json
 */
async function retrieveABuiltInPolicyDefinition() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const policyDefinitionName = "7433c107-6db4-4ad1-b57a-a76dce0154a1";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.policyDefinitions.getBuiltIn(policyDefinitionName);
  console.log(result);
}
