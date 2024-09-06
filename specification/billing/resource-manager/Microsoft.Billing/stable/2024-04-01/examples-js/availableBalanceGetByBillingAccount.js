const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The Available Credit or Payment on Account Balance for a billing account. The credit balance can be used to settle due or past due invoices and is supported for billing accounts with agreement type Microsoft Customer Agreement. The payment on account balance is supported for billing accounts with agreement type Microsoft Customer Agreement or Microsoft Online Services Program.
 *
 * @summary The Available Credit or Payment on Account Balance for a billing account. The credit balance can be used to settle due or past due invoices and is supported for billing accounts with agreement type Microsoft Customer Agreement. The payment on account balance is supported for billing accounts with agreement type Microsoft Customer Agreement or Microsoft Online Services Program.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/availableBalanceGetByBillingAccount.json
 */
async function availableBalanceGetByBillingAccount() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.availableBalances.getByBillingAccount(billingAccountName);
  console.log(result);
}
