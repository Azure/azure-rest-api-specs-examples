const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a department by ID. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
 *
 * @summary Gets a department by ID. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/departmentGet.json
 */
async function departmentGet() {
  const billingAccountName = "456598";
  const departmentName = "164821";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.departments.get(billingAccountName, departmentName);
  console.log(result);
}
