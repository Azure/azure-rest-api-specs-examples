const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an invoice section. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 *
 * @summary Creates or updates an invoice section. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutInvoiceSection.json
 */
async function putInvoiceSection() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const invoiceSectionName = "{invoiceSectionName}";
  const parameters = {
    displayName: "invoiceSection1",
    labels: { costCategory: "Support", pcCode: "A123456" },
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.invoiceSections.beginCreateOrUpdateAndWait(
    billingAccountName,
    billingProfileName,
    invoiceSectionName,
    parameters
  );
  console.log(result);
}

putInvoiceSection().catch(console.error);
