const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a role assignment by scope and name.
 *
 * @summary Get a role assignment by scope and name.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_Get.json
 */
async function getRoleAssignmentByScopeAndName() {
  const scope = "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2";
  const roleAssignmentName = "b0f43c54-e787-4862-89b1-a653fa9cf747";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.roleAssignments.get(scope, roleAssignmentName);
  console.log(result);
}
