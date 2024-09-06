const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides a list of check access response objects for an invoice section.
 *
 * @summary Provides a list of check access response objects for an invoice section.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/checkAccessByInvoiceSection.json
 */
async function checkAccessByInvoiceSection() {
  const billingAccountName =
    "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const billingProfileName = "xxxx-xxxx-xxx-xxx";
  const invoiceSectionName = "Q7GV-UUVA-PJA-TGB";
  const parameters = {
    actions: [
      "Microsoft.Billing/billingAccounts/read",
      "Microsoft.Subscription/subscriptions/write",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingPermissions.checkAccessByInvoiceSection(
    billingAccountName,
    billingProfileName,
    invoiceSectionName,
    parameters,
  );
  console.log(result);
}
