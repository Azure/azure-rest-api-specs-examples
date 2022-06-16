const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Kusto cluster principalAssignments.
 *
 * @summary Lists all Kusto cluster principalAssignments.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoClusterPrincipalAssignmentsList.json
 */
async function kustoPrincipalAssignmentsList() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.clusterPrincipalAssignments.list(resourceGroupName, clusterName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

kustoPrincipalAssignmentsList().catch(console.error);
