const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the Front Doors within an Azure subscription.
 *
 * @summary Lists all of the Front Doors within an Azure subscription.
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorListAll.json
 */
async function listAllFrontDoors() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.frontDoors.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllFrontDoors().catch(console.error);
