const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Result of the request to list REST API operations
 *
 * @summary Result of the request to list REST API operations
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ProviderOperations_List.json
 */
async function providerOperationsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.providerOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

providerOperationsList().catch(console.error);
