const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all linked storage accounts associated with the specified workspace, storage accounts will be sorted by their data source type.
 *
 * @summary Gets all linked storage accounts associated with the specified workspace, storage accounts will be sorted by their data source type.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/LinkedStorageAccountsListByWorkspace.json
 */
async function getsListOfLinkedStorageAccountsOnAWorkspace() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "mms-eus";
  const workspaceName = "testLinkStorageAccountsWS";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.linkedStorageAccounts.listByWorkspace(
    resourceGroupName,
    workspaceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsListOfLinkedStorageAccountsOnAWorkspace().catch(console.error);
