const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the subscriptions for an enrollment account. The operation is supported for billing accounts with agreement type Enterprise Agreement.
 *
 * @summary Lists the subscriptions for an enrollment account. The operation is supported for billing accounts with agreement type Enterprise Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionsListByEnrollmentAccount.json
 */
async function billingSubscriptionsListByEnrollmentAccount() {
  const billingAccountName = "6564892";
  const enrollmentAccountName = "172988";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.billingSubscriptions.listByEnrollmentAccount(
    billingAccountName,
    enrollmentAccountName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
