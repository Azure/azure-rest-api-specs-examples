const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a role assignment for the caller on an invoice section. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement.
 *
 * @summary Deletes a role assignment for the caller on an invoice section. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionRoleAssignmentDelete.json
 */
async function invoiceSectionRoleAssignmentDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const invoiceSectionName = "{invoiceSectionName}";
  const billingRoleAssignmentName = "{billingRoleAssignmentName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingRoleAssignments.deleteByInvoiceSection(
    billingAccountName,
    billingProfileName,
    invoiceSectionName,
    billingRoleAssignmentName
  );
  console.log(result);
}

invoiceSectionRoleAssignmentDelete().catch(console.error);
