const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks availability and correctness of the name for a scheduled action.
 *
 * @summary Checks availability and correctness of the name for a scheduled action.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/checkNameAvailability-private-scheduledAction.json
 */
async function scheduledActionCheckNameAvailability() {
  const checkNameAvailabilityRequest = {
    name: "testName",
    type: "Microsoft.CostManagement/ScheduledActions",
  };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.scheduledActions.checkNameAvailability(checkNameAvailabilityRequest);
  console.log(result);
}
