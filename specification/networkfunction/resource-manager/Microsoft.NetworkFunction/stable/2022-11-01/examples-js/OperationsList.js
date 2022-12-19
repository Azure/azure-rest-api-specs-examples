const { AzureTrafficCollectorClient } = require("@azure/arm-networkfunction");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available NetworkFunction Rest API operations.
 *
 * @summary Lists all of the available NetworkFunction Rest API operations.
 * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/OperationsList.json
 */
async function operationsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new AzureTrafficCollectorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkFunction.listOperations()) {
    resArray.push(item);
  }
  console.log(resArray);
}

operationsList().catch(console.error);
