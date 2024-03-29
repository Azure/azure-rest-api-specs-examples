const { DevSpacesManagementClient } = require("@azure/arm-devspaces");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the Azure Dev Spaces Controllers with their properties in the subscription.
 *
 * @summary Lists all the Azure Dev Spaces Controllers with their properties in the subscription.
 * x-ms-original-file: specification/devspaces/resource-manager/Microsoft.DevSpaces/stable/2019-04-01/examples/ControllersList_example.json
 */
async function controllersList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new DevSpacesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.controllers.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

controllersList().catch(console.error);
