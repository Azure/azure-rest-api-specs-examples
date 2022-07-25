const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the database principal assignment is valid and is not already in use.
 *
 * @summary Checks that the database principal assignment is valid and is not already in use.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasePrincipalAssignmentsCheckNameAvailability.json
 */
async function kustoPoolDatabasePrincipalAssignmentsCheckNameAvailability() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const workspaceName = "synapseWorkspaceName";
  const kustoPoolName = "kustoclusterrptest4";
  const databaseName = "Kustodatabase8";
  const resourceGroupName = "kustorptest";
  const principalAssignmentName = {
    name: "kustoprincipal1",
    type: "Microsoft.Synapse/workspaces/kustoPools/databases/principalAssignments",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPoolDatabasePrincipalAssignments.checkNameAvailability(
    workspaceName,
    kustoPoolName,
    databaseName,
    resourceGroupName,
    principalAssignmentName
  );
  console.log(result);
}

kustoPoolDatabasePrincipalAssignmentsCheckNameAvailability().catch(console.error);
