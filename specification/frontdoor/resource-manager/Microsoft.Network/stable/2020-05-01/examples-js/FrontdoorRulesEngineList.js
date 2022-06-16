const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the Rules Engine Configurations within a Front Door.
 *
 * @summary Lists all of the Rules Engine Configurations within a Front Door.
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorRulesEngineList.json
 */
async function listRulesEngineConfigurationsInAFrontDoor() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const frontDoorName = "frontDoor1";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.rulesEngines.listByFrontDoor(resourceGroupName, frontDoorName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listRulesEngineConfigurationsInAFrontDoor().catch(console.error);
