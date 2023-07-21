const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified deny assignment.
 *
 * @summary Get the specified deny assignment.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/GetDenyAssignmentByNameId.json
 */
async function getDenyAssignmentByName() {
  const scope = "subscriptions/subId/resourcegroups/rgname";
  const denyAssignmentId = "denyAssignmentId";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.denyAssignments.get(scope, denyAssignmentId);
  console.log(result);
}
