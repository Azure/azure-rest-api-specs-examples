const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of private endpoint connections.
 *
 * @summary Returns the list of private endpoint connections.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoPrivateEndpointConnectionsList.json
 */
async function kustoPrivateEndpointConnectionsList() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.list(resourceGroupName, clusterName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

kustoPrivateEndpointConnectionsList().catch(console.error);
