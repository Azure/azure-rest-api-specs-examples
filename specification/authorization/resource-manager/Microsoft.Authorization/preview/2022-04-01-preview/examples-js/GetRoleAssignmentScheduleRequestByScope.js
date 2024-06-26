const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets role assignment schedule requests for a scope.
 *
 * @summary Gets role assignment schedule requests for a scope.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-04-01-preview/examples/GetRoleAssignmentScheduleRequestByScope.json
 */
async function getRoleAssignmentScheduleRequestByScope() {
  const scope =
    "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f";
  const filter = "assignedTo('A3BB8764-CB92-4276-9D2A-CA1E895E55EA')";
  const options = {
    filter,
  };
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.roleAssignmentScheduleRequests.listForScope(scope, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
