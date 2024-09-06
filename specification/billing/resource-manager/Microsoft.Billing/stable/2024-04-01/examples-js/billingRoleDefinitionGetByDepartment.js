const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the definition for a role on a department. The operation is supported for billing accounts with agreement type Enterprise Agreement.
 *
 * @summary Gets the definition for a role on a department. The operation is supported for billing accounts with agreement type Enterprise Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleDefinitionGetByDepartment.json
 */
async function billingRoleDefinitionGetByDepartment() {
  const billingAccountName = "123456";
  const departmentName = "7368531";
  const roleDefinitionName = "50000000-aaaa-bbbb-cccc-100000000000";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingRoleDefinitionOperations.getByDepartment(
    billingAccountName,
    departmentName,
    roleDefinitionName,
  );
  console.log(result);
}
