const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to list all exports at the given scope.
 *
 * @summary the operation to list all exports at the given scope.
 * x-ms-original-file: 2025-03-01/ExportsGetByBillingAccount.json
 */
async function exportsGetByBillingAccount() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.list("providers/Microsoft.Billing/billingAccounts/123456");
  console.log(result);
}
