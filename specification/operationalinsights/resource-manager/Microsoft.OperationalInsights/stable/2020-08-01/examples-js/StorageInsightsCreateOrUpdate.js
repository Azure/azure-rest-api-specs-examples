const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a storage insight.
 *
 * @summary Create or update a storage insight.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/StorageInsightsCreateOrUpdate.json
 */
async function storageInsightsCreate() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "OIAutoRest5123";
  const workspaceName = "aztest5048";
  const storageInsightName = "AzTestSI1110";
  const parameters = {
    containers: ["wad-iis-logfiles"],
    storageAccount: {
      id: "/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/OIAutoRest6987/providers/microsoft.storage/storageaccounts/AzTestFakeSA9945",
      key: "1234",
    },
    tables: ["WADWindowsEventLogsTable", "LinuxSyslogVer2v0"],
  };
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.storageInsightConfigs.createOrUpdate(
    resourceGroupName,
    workspaceName,
    storageInsightName,
    parameters
  );
  console.log(result);
}

storageInsightsCreate().catch(console.error);
