const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the balances for a scope by billingAccountId. Balances are available via this API only for May 1, 2014 or later.
 *
 * @summary Gets the balances for a scope by billingAccountId. Balances are available via this API only for May 1, 2014 or later.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/BalancesByBillingAccount.json
 */
async function balances() {
  const subscriptionId =
    process.env["CONSUMPTION_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const billingAccountId = "123456";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const result = await client.balances.getByBillingAccount(billingAccountId);
  console.log(result);
}
