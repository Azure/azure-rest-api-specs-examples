const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a Log Analytics workspace table properties.
 *
 * @summary Updates a Log Analytics workspace table properties.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/TablesCreate.json
 */
async function tablesSet() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "oiautorest6685";
  const workspaceName = "oiautorest6685";
  const tableName = "table1";
  const parameters = {
    isTroubleshootEnabled: false,
    retentionInDays: 40,
  };
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.tables.create(
    resourceGroupName,
    workspaceName,
    tableName,
    parameters
  );
  console.log(result);
}

tablesSet().catch(console.error);
