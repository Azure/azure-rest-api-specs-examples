const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to list all exports at the given scope.
 *
 * @summary The operation to list all exports at the given scope.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportsGetByBillingAccount.json
 */
async function exportsGetByBillingAccount() {
  const scope = "providers/Microsoft.Billing/billingAccounts/123456";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.list(scope);
  console.log(result);
}
