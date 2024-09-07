const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the enrollment accounts for a billing account. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
 *
 * @summary Lists the enrollment accounts for a billing account. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/enrollmentAccountsListByBillingAccount.json
 */
async function enrollmentAccountsListByBillingAccount() {
  const billingAccountName = "6564892";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.enrollmentAccounts.listByBillingAccount(billingAccountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
