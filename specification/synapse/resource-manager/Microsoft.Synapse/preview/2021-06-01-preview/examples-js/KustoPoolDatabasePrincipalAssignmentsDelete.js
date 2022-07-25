const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Kusto pool principalAssignment.
 *
 * @summary Deletes a Kusto pool principalAssignment.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasePrincipalAssignmentsDelete.json
 */
async function kustoPoolDatabasePrincipalAssignmentsDelete() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const workspaceName = "synapseWorkspaceName";
  const kustoPoolName = "kustoclusterrptest4";
  const databaseName = "Kustodatabase8";
  const principalAssignmentName = "kustoprincipal1";
  const resourceGroupName = "kustorptest";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPoolDatabasePrincipalAssignments.beginDeleteAndWait(
    workspaceName,
    kustoPoolName,
    databaseName,
    principalAssignmentName,
    resourceGroupName
  );
  console.log(result);
}

kustoPoolDatabasePrincipalAssignmentsDelete().catch(console.error);
