const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Checks that the principal assignment name is valid and is not already in use.
 *
 * @summary Checks that the principal assignment name is valid and is not already in use.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoClusterPrincipalAssignmentsCheckNameAvailability.json
 */
async function kustoClusterPrincipalAssignmentsCheckNameAvailability() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["KUSTO_RESOURCE_GROUP"] || "kustorptest";
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
    principalAssignmentName,
  );
  console.log(result);
}
