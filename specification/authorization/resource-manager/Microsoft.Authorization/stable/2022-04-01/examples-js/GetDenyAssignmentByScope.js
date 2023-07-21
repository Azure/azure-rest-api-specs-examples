const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets deny assignments for a scope.
 *
 * @summary Gets deny assignments for a scope.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/GetDenyAssignmentByScope.json
 */
async function listDenyAssignmentsForScope() {
  const scope = "subscriptions/subId";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.denyAssignments.listForScope(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
