const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks availability and correctness of the name for a scheduled action within the given scope.
 *
 * @summary Checks availability and correctness of the name for a scheduled action within the given scope.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/checkNameAvailability-shared-scheduledAction.json
 */
async function scheduledActionCheckNameAvailabilityByScope() {
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000";
  const checkNameAvailabilityRequest = {
    name: "testName",
    type: "Microsoft.CostManagement/ScheduledActions",
  };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.scheduledActions.checkNameAvailabilityByScope(
    scope,
    checkNameAvailabilityRequest
  );
  console.log(result);
}
