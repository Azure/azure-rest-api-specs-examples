const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a role assignment on a enrollment Account. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
 *
 * @summary Deletes a role assignment on a enrollment Account. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentDeleteByEnrollmentAccount.json
 */
async function billingRoleAssignmentDeleteByEnrollmentAccount() {
  const billingAccountName = "8608480";
  const enrollmentAccountName = "123456";
  const billingRoleAssignmentName = "9dfd08c2-62a3-4d47-85bd-1cdba1408402";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingRoleAssignments.deleteByEnrollmentAccount(
    billingAccountName,
    enrollmentAccountName,
    billingRoleAssignmentName,
  );
  console.log(result);
}
