const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides the aggregate cost of a management group and all child management groups by current billing period.
 *
 * @summary Provides the aggregate cost of a management group and all child management groups by current billing period.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/AggregatedCostByManagementGroup.json
 */
async function aggregatedCostByManagementGroup() {
  const subscriptionId =
    process.env["CONSUMPTION_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const managementGroupId = "managementGroupForTest";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const result = await client.aggregatedCost.getByManagementGroup(managementGroupId);
  console.log(result);
}
