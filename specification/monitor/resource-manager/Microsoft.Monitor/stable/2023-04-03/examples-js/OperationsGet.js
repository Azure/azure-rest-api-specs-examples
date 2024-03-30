const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists available Operations for this Resource Provider
 *
 * @summary Lists available Operations for this Resource Provider
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Monitor/stable/2023-04-03/examples/OperationsGet.json
 */
async function getOperationsList() {
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const resArray = new Array();
  for await (let item of client.monitorOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
