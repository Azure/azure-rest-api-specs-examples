const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Front Door with the specified Front Door name under the specified subscription and resource group.
 *
 * @summary Gets a Front Door with the specified Front Door name under the specified subscription and resource group.
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorGet.json
 */
async function getFrontDoor() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const frontDoorName = "frontDoor1";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.frontDoors.get(resourceGroupName, frontDoorName);
  console.log(result);
}

getFrontDoor().catch(console.error);
