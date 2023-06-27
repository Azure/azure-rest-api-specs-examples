const { HelpRP } = require("@azure/arm-selfhelp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of operations.
 *
 * @summary Returns list of operations.
 * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/stable/2023-06-01/examples/ListOperations.json
 */
async function listAllOperations() {
  const credential = new DefaultAzureCredential();
  const client = new HelpRP(credential);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
