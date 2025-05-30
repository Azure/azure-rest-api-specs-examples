const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation lists all the policy definition versions for all policy definitions within a subscription.
 *
 * @summary This operation lists all the policy definition versions for all policy definitions within a subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2024-05-01/examples/listAllPolicyDefinitionVersions.json
 */
async function listAllPolicyDefinitionVersionsAtSubscription() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.policyDefinitionVersions.listAll();
  console.log(result);
}
