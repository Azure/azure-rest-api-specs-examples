const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a data export instance.
 *
 * @summary Gets a data export instance.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataExportGet.json
 */
async function dataExportGet() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "RgTest1";
  const workspaceName = "DeWnTest1234";
  const dataExportName = "export1";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.dataExports.get(resourceGroupName, workspaceName, dataExportName);
  console.log(result);
}

dataExportGet().catch(console.error);
