const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the definition for a role on an enrollment account. The operation is supported for billing accounts with agreement type Enterprise Agreement.
 *
 * @summary Gets the definition for a role on an enrollment account. The operation is supported for billing accounts with agreement type Enterprise Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleDefinitionGetByEnrollmentAccount.json
 */
async function billingRoleDefinitionGetByEnrollmentAccount() {
  const billingAccountName = "123456";
  const enrollmentAccountName = "4568789";
  const roleDefinitionName = "50000000-aaaa-bbbb-cccc-100000000000";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingRoleDefinitionOperations.getByEnrollmentAccount(
    billingAccountName,
    enrollmentAccountName,
    roleDefinitionName,
  );
  console.log(result);
}
