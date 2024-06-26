const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The available credit balance for a billing profile. This is the balance that can be used for pay now to settle due or past due invoices. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 *
 * @summary The available credit balance for a billing profile. This is the balance that can be used for pay now to settle due or past due invoices. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/AvailableBalanceByBillingProfile.json
 */
async function availableBalanceByBillingProfile() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.availableBalances.get(billingAccountName, billingProfileName);
  console.log(result);
}

availableBalanceByBillingProfile().catch(console.error);
