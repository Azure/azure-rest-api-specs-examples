const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Kusto clusters within a subscription.
 *
 * @summary Lists all Kusto clusters within a subscription.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoClustersList.json
 */
async function kustoClustersList() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.clusters.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

kustoClustersList().catch(console.error);
