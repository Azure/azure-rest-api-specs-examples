const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the storage accounts which are in use by related reports
 *
 * @summary List the storage accounts which are in use by related reports
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/ListInUseStorageAccountsWithoutSubscriptions.json
 */
async function listInUseStorageAccountsWithoutSubscriptions() {
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.providerActions.listInUseStorageAccounts({});
  console.log(result);
}
