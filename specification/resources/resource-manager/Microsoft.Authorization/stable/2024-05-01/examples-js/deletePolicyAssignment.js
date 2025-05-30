const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation deletes a policy assignment, given its name and the scope it was created in. The scope of a policy assignment is the part of its ID preceding '/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}'.
 *
 * @summary This operation deletes a policy assignment, given its name and the scope it was created in. The scope of a policy assignment is the part of its ID preceding '/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}'.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2024-05-01/examples/deletePolicyAssignment.json
 */
async function deleteAPolicyAssignment() {
  const scope = "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const policyAssignmentName = "EnforceNaming";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyAssignments.delete(scope, policyAssignmentName);
  console.log(result);
}
