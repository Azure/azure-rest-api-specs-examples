const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the principal assignment name is valid and is not already in use.
 *
 * @summary Checks that the principal assignment name is valid and is not already in use.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoClusterPrincipalAssignmentsCheckNameAvailability.json
 */
async function kustoClusterPrincipalAssignmentsCheckNameAvailability() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const principalAssignmentName = {
    name: "kustoprincipal1",
    type: "Microsoft.Kusto/clusters/principalAssignments",
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusterPrincipalAssignments.checkNameAvailability(
    resourceGroupName,
    clusterName,
    principalAssignmentName
  );
  console.log(result);
}

kustoClusterPrincipalAssignmentsCheckNameAvailability().catch(console.error);
