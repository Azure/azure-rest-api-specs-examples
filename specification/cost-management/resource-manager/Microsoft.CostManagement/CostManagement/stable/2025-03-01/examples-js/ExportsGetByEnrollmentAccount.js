const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to list all exports at the given scope.
 *
 * @summary the operation to list all exports at the given scope.
 * x-ms-original-file: 2025-03-01/ExportsGetByEnrollmentAccount.json
 */
async function exportsGetByEnrollmentAccount() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.list(
    "providers/Microsoft.Billing/billingAccounts/100/enrollmentAccounts/456",
  );
  console.log(result);
}
