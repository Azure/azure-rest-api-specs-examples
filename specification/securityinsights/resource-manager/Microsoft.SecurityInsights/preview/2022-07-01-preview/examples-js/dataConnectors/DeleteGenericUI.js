const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the data connector.
 *
 * @summary Delete the data connector.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/dataConnectors/DeleteGenericUI.json
 */
async function deleteAGenericUiDataConnector() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const dataConnectorId = "316ec55e-7138-4d63-ab18-90c8a60fd1c8";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectors.delete(
    resourceGroupName,
    workspaceName,
    dataConnectorId
  );
  console.log(result);
}

deleteAGenericUiDataConnector().catch(console.error);
