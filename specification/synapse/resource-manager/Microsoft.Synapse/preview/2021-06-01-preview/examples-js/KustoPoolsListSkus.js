const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists eligible SKUs for Kusto Pool resource.
 *
 * @summary Lists eligible SKUs for Kusto Pool resource.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsListSkus.json
 */
async function kustoPoolsListSkus() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.kustoPools.listSkus()) {
    resArray.push(item);
  }
  console.log(resArray);
}

kustoPoolsListSkus().catch(console.error);
