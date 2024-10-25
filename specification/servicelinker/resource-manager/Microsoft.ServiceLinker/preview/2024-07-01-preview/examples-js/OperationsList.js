const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the available ServiceLinker REST API operations.
 *
 * @summary Lists the available ServiceLinker REST API operations.
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/OperationsList.json
 */
async function getConfiguration() {
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
