const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all linked storage account of a specific data source type associated with the specified workspace.
 *
 * @summary Gets all linked storage account of a specific data source type associated with the specified workspace.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/LinkedStorageAccountsGet.json
 */
async function linkedStorageAccountsGet() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "mms-eus";
  const workspaceName = "testLinkStorageAccountsWS";
  const dataSourceType = "CustomLogs";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.linkedStorageAccounts.get(
    resourceGroupName,
    workspaceName,
    dataSourceType
  );
  console.log(result);
}

linkedStorageAccountsGet().catch(console.error);
