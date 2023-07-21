const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a deny assignment by ID.
 *
 * @summary Gets a deny assignment by ID.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/GetDenyAssignmentById.json
 */
async function getDenyAssignmentById() {
  const denyAssignmentId =
    "subscriptions/subId/resourcegroups/rgname/providers/Microsoft.Authorization/denyAssignments/daId";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.denyAssignments.getById(denyAssignmentId);
  console.log(result);
}
