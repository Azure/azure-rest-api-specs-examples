const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Cancels a pending role eligibility schedule request.
 *
 * @summary Cancels a pending role eligibility schedule request.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-04-01-preview/examples/CancelRoleEligibilityScheduleRequestByName.json
 */
async function cancelRoleEligibilityScheduleRequestByName() {
  const scope =
    "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f";
  const roleEligibilityScheduleRequestName = "64caffb6-55c0-4deb-a585-68e948ea1ad6";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.roleEligibilityScheduleRequests.cancel(
    scope,
    roleEligibilityScheduleRequestName
  );
  console.log(result);
}
