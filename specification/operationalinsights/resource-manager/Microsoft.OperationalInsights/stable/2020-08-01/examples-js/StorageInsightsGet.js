const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a storage insight instance.
 *
 * @summary Gets a storage insight instance.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/StorageInsightsGet.json
 */
async function storageInsightsGet() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "OIAutoRest5123";
  const workspaceName = "aztest5048";
  const storageInsightName = "AzTestSI1110";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.storageInsightConfigs.get(
    resourceGroupName,
    workspaceName,
    storageInsightName
  );
  console.log(result);
}

storageInsightsGet().catch(console.error);
