const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the integration runtime node.
 *
 * @summary Delete the integration runtime node.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/DeleteIntegrationRuntimeNode.json
 */
async function deleteTheIntegrationRuntimeNode() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const sqlMigrationServiceName = "service1";
  const parameters = {
    integrationRuntimeName: "IRName",
    nodeName: "nodeName",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.sqlMigrationServices.deleteNode(
    resourceGroupName,
    sqlMigrationServiceName,
    parameters
  );
  console.log(result);
}

deleteTheIntegrationRuntimeNode().catch(console.error);
