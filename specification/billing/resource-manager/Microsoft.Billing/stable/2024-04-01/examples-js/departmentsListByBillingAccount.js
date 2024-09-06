const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the departments that a user has access to. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
 *
 * @summary Lists the departments that a user has access to. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/departmentsListByBillingAccount.json
 */
async function departmentsListByBillingAccount() {
  const billingAccountName = "456598";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.departments.listByBillingAccount(billingAccountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
