const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to this operation retrieves a single policy assignment, given its name and the scope it was created at.
 *
 * @summary this operation retrieves a single policy assignment, given its name and the scope it was created at.
 * x-ms-original-file: 2025-03-01/getPolicyAssignmentWithUserAssignedIdentity.json
 */
async function retrieveAPolicyAssignmentWithAUserAssignedIdentity() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyAssignments.get(
    "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2",
    "EnforceNaming",
  );
  console.log(result);
}
