const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Log Analytics workspace table.
 *
 * @summary Gets a Log Analytics workspace table.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/TablesGet.json
 */
async function tablesGet() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "oiautorest6685";
  const workspaceName = "oiautorest6685";
  const tableName = "table1_SRCH";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.tables.get(resourceGroupName, workspaceName, tableName);
  console.log(result);
}

tablesGet().catch(console.error);
