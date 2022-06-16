const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a storageInsightsConfigs resource
 *
 * @summary Deletes a storageInsightsConfigs resource
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/StorageInsightsDelete.json
 */
async function storageInsightsDelete() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "OIAutoRest5123";
  const workspaceName = "aztest5048";
  const storageInsightName = "AzTestSI1110";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.storageInsightConfigs.delete(
    resourceGroupName,
    workspaceName,
    storageInsightName
  );
  console.log(result);
}

storageInsightsDelete().catch(console.error);
