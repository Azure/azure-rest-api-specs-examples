const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all role assignments that apply to a scope.
 *
 * @summary List all role assignments that apply to a scope.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_ListForScope.json
 */
async function listRoleAssignmentsForScope() {
  const subscriptionId =
    process.env["AUTHORIZATION_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.roleAssignments.listForScope(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
