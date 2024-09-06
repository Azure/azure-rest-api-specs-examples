const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an invoice by billing account name and ID. The operation is supported for all billing account types.
 *
 * @summary Gets an invoice by billing account name and ID. The operation is supported for all billing account types.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoicesGetByBillingAccount.json
 */
async function invoicesGetByBillingAccount() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const invoiceName = "G123456789";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.invoices.getByBillingAccount(billingAccountName, invoiceName);
  console.log(result);
}
