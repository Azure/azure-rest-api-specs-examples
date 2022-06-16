const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists eligible SKUs for Kusto resource provider.
 *
 * @summary Lists eligible SKUs for Kusto resource provider.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoClustersListSkus.json
 */
async function kustoClustersListSkus() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.clusters.listSkus()) {
    resArray.push(item);
  }
  console.log(resArray);
}

kustoClustersListSkus().catch(console.error);
