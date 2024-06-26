const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides the aggregate cost of a management group and all child management groups by specified billing period
 *
 * @summary Provides the aggregate cost of a management group and all child management groups by specified billing period
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/AggregatedCostForBillingPeriodByManagementGroup.json
 */
async function aggregatedCostListForBillingPeriodByManagementGroup() {
  const subscriptionId =
    process.env["CONSUMPTION_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const managementGroupId = "managementGroupForTest";
  const billingPeriodName = "201807";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const result = await client.aggregatedCost.getForBillingPeriodByManagementGroup(
    managementGroupId,
    billingPeriodName
  );
  console.log(result);
}
