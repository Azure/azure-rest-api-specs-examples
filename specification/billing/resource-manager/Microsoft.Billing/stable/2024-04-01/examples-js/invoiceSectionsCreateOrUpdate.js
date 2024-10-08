const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an invoice section. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 *
 * @summary Creates or updates an invoice section. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoiceSectionsCreateOrUpdate.json
 */
async function invoiceSectionsCreateOrUpdate() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const billingProfileName = "xxxx-xxxx-xxx-xxx";
  const invoiceSectionName = "invoice-section-1";
  const parameters = {
    properties: {
      displayName: "Invoice Section 1",
      tags: { costCategory: "Support", pcCode: "A123456" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.invoiceSections.beginCreateOrUpdateAndWait(
    billingAccountName,
    billingProfileName,
    invoiceSectionName,
    parameters,
  );
  console.log(result);
}
