const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Kusto pool database principalAssignments.
 *
 * @summary Lists all Kusto pool database principalAssignments.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasePrincipalAssignmentsList.json
 */
async function kustoPoolDatabasePrincipalAssignmentsList() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const workspaceName = "synapseWorkspaceName";
  const kustoPoolName = "kustoclusterrptest4";
  const databaseName = "Kustodatabase8";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "kustorptest";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.kustoPoolDatabasePrincipalAssignments.list(
    workspaceName,
    kustoPoolName,
    databaseName,
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
