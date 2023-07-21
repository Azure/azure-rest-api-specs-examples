const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified role assignment schedule request.
 *
 * @summary Get the specified role assignment schedule request.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-04-01-preview/examples/GetRoleAssignmentScheduleRequestByName.json
 */
async function getRoleAssignmentScheduleRequestByName() {
  const scope =
    "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f";
  const roleAssignmentScheduleRequestName = "fea7a502-9a96-4806-a26f-eee560e52045";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.roleAssignmentScheduleRequests.get(
    scope,
    roleAssignmentScheduleRequestName
  );
  console.log(result);
}
