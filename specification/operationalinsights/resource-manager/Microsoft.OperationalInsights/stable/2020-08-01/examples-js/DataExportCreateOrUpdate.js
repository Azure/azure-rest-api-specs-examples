const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a data export.
 *
 * @summary Create or update a data export.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataExportCreateOrUpdate.json
 */
async function dataExportCreate() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "RgTest1";
  const workspaceName = "DeWnTest1234";
  const dataExportName = "export1";
  const parameters = {
    resourceId:
      "/subscriptions/192b9f85-a39a-4276-b96d-d5cd351703f9/resourceGroups/OIAutoRest1234/providers/Microsoft.EventHub/namespaces/test",
    tableNames: ["Heartbeat"],
  };
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.dataExports.createOrUpdate(
    resourceGroupName,
    workspaceName,
    dataExportName,
    parameters
  );
  console.log(result);
}

dataExportCreate().catch(console.error);
