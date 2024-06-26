const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a role assignment for the caller on an invoice section. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement.
 *
 * @summary Gets a role assignment for the caller on an invoice section. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionRoleAssignment.json
 */
async function invoiceSectionRoleAssignment() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const invoiceSectionName = "{invoiceSectionName}";
  const billingRoleAssignmentName = "{billingRoleAssignmentName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingRoleAssignments.getByInvoiceSection(
    billingAccountName,
    billingProfileName,
    invoiceSectionName,
    billingRoleAssignmentName
  );
  console.log(result);
}

invoiceSectionRoleAssignment().catch(console.error);
