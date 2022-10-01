const { ContainerAppsAPIClient } = require("@azure/arm-app");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available RP operations.
 *
 * @summary Lists all of the available RP operations.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/Operations_List.json
 */
async function listAllOperations() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllOperations().catch(console.error);
