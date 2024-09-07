const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the subscriptions for a billing account.
 *
 * @summary Lists the subscriptions for a billing account.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionsListByBillingAccount.json
 */
async function billingSubscriptionsListByBillingAccount() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const includeDeleted = false;
  const includeTenantSubscriptions = false;
  const options = {
    includeDeleted,
    includeTenantSubscriptions,
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.billingSubscriptions.listByBillingAccount(
    billingAccountName,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
