const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves a single policy assignment, given its name and the scope it was created at.
 *
 * @summary This operation retrieves a single policy assignment, given its name and the scope it was created at.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2024-05-01/examples/getPolicyAssignmentWithOverrides.json
 */
async function retrieveAPolicyAssignmentWithOverrides() {
  const scope = "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const policyAssignmentName = "CostManagement";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyAssignments.get(scope, policyAssignmentName);
  console.log(result);
}
