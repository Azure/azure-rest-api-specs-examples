const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a data source.
 *
 * @summary Create or update a data source.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataSourcesCreate.json
 */
async function dataSourcesCreate() {
  const subscriptionId =
    process.env["OPERATIONALINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = process.env["OPERATIONALINSIGHTS_RESOURCE_GROUP"] || "OIAutoRest5123";
  const workspaceName = "AzTest9724";
  const dataSourceName = "AzTestDS774";
  const parameters = {
    kind: "AzureActivityLog",
    properties: {
      LinkedResourceId:
        "/subscriptions/00000000-0000-0000-0000-00000000000/providers/microsoft.insights/eventtypes/management",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.dataSources.createOrUpdate(
    resourceGroupName,
    workspaceName,
    dataSourceName,
    parameters
  );
  console.log(result);
}
