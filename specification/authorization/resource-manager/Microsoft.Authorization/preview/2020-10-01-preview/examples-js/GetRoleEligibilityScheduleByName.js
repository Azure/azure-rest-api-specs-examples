const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified role eligibility schedule for a resource scope
 *
 * @summary Get the specified role eligibility schedule for a resource scope
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2020-10-01-preview/examples/GetRoleEligibilityScheduleByName.json
 */
async function getRoleEligibilityScheduleByName() {
  const scope =
    "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f";
  const roleEligibilityScheduleName = "b1477448-2cc6-4ceb-93b4-54a202a89413";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.roleEligibilitySchedules.get(scope, roleEligibilityScheduleName);
  console.log(result);
}
