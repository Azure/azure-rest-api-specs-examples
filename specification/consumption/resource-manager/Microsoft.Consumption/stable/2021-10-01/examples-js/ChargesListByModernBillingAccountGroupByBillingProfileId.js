const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the charges based for the defined scope.
 *
 * @summary Lists the charges based for the defined scope.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ChargesListByModernBillingAccountGroupByBillingProfileId.json
 */
async function chargesListByBillingAccountGroupByBillingProfileIdModern() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "providers/Microsoft.Billing/billingAccounts/1234:56789";
  const startDate = "2019-09-01";
  const endDate = "2019-09-30";
  const apply = "groupby((properties/billingProfileId))";
  const options = { startDate, endDate, apply };
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const result = await client.charges.list(scope, options);
  console.log(result);
}

chargesListByBillingAccountGroupByBillingProfileIdModern().catch(console.error);
