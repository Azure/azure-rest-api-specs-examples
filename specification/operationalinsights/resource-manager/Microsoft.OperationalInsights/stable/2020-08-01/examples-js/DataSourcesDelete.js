const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a data source instance.
 *
 * @summary Deletes a data source instance.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataSourcesDelete.json
 */
async function dataSourcesDelete() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "OIAutoRest5123";
  const workspaceName = "AzTest9724";
  const dataSourceName = "AzTestDS774";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.dataSources.delete(resourceGroupName, workspaceName, dataSourceName);
  console.log(result);
}

dataSourcesDelete().catch(console.error);
