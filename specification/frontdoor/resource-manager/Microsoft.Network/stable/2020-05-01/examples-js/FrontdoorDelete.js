const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing Front Door with the specified parameters.
 *
 * @summary Deletes an existing Front Door with the specified parameters.
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorDelete.json
 */
async function deleteFrontDoor() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const frontDoorName = "frontDoor1";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.frontDoors.beginDeleteAndWait(resourceGroupName, frontDoorName);
  console.log(result);
}

deleteFrontDoor().catch(console.error);
