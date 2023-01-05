const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the marketplaces for a scope at the defined scope. Marketplaces are available via this API only for May 1, 2014 or later.
 *
 * @summary Lists the marketplaces for a scope at the defined scope. Marketplaces are available via this API only for May 1, 2014 or later.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/MarketplacesByEnrollmentAccounts_ListByBillingPeriod.json
 */
async function enrollmentAccountMarketplacesListForBillingPeriod() {
  const subscriptionId =
    process.env["CONSUMPTION_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope = "providers/Microsoft.Billing/enrollmentAccounts/123456";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.marketplaces.list(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
