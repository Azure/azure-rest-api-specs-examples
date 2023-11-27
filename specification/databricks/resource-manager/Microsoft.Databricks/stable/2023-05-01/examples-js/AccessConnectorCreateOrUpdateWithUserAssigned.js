const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates azure databricks accessConnector.
 *
 * @summary Creates or updates azure databricks accessConnector.
 * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-05-01/examples/AccessConnectorCreateOrUpdateWithUserAssigned.json
 */
async function createAnAzureDatabricksAccessConnectorWithUserAssignedIdentity() {
  const subscriptionId = process.env["DATABRICKS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["DATABRICKS_RESOURCE_GROUP"] || "rg";
  const connectorName = "myAccessConnector";
  const parameters = { location: "westus" };
  const credential = new DefaultAzureCredential();
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.accessConnectors.beginCreateOrUpdateAndWait(
    resourceGroupName,
    connectorName,
    parameters
  );
  console.log(result);
}
