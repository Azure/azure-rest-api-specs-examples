const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a datasource instance.
 *
 * @summary Gets a datasource instance.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataSourcesGet.json
 */
async function dataSourcesGet() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "OIAutoRest5123";
  const workspaceName = "AzTest9724";
  const dataSourceName = "AzTestDS774";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.dataSources.get(resourceGroupName, workspaceName, dataSourceName);
  console.log(result);
}

dataSourcesGet().catch(console.error);
