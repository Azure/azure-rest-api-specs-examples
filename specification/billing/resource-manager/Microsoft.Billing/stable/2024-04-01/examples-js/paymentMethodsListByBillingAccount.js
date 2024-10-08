const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the payment methods available for a billing account. Along with the payment methods owned by the caller, these payment methods can be attached to a billing profile to make payments. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 *
 * @summary Lists the payment methods available for a billing account. Along with the payment methods owned by the caller, these payment methods can be attached to a billing profile to make payments. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsListByBillingAccount.json
 */
async function paymentMethodsListByBillingAccount() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.paymentMethods.listByBillingAccount(billingAccountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
