const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves a list of all the policy set definition versions for the given policy set definition.
 *
 * @summary This operation retrieves a list of all the policy set definition versions for the given policy set definition.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2024-05-01/examples/listPolicySetDefinitionVersions.json
 */
async function listPolicySetDefinitions() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const policySetDefinitionName = "CostManagement";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.policySetDefinitionVersions.list(policySetDefinitionName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
