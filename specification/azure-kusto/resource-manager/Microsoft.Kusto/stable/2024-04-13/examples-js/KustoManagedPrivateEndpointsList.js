const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Returns the list of managed private endpoints.
 *
 * @summary Returns the list of managed private endpoints.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoManagedPrivateEndpointsList.json
 */
async function kustoManagedPrivateEndpointsList() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["KUSTO_RESOURCE_GROUP"] || "kustorptest";
  const clusterName = "kustoCluster";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedPrivateEndpoints.list(resourceGroupName, clusterName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
