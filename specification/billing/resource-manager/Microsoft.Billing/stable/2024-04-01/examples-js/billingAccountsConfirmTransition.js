const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the transition details for a billing account that has transitioned from agreement type Microsoft Online Services Program to agreement type Microsoft Customer Agreement.
 *
 * @summary Gets the transition details for a billing account that has transitioned from agreement type Microsoft Online Services Program to agreement type Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingAccountsConfirmTransition.json
 */
async function billingAccountsConfirmTransition() {
  const billingAccountName =
    "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingAccounts.confirmTransition(billingAccountName);
  console.log(result);
}
