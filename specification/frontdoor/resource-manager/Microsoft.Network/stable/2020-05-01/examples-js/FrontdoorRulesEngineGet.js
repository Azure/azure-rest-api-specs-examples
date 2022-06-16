const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Rules Engine Configuration with the specified name within the specified Front Door.
 *
 * @summary Gets a Rules Engine Configuration with the specified name within the specified Front Door.
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorRulesEngineGet.json
 */
async function getRulesEngineConfiguration() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const frontDoorName = "frontDoor1";
  const rulesEngineName = "rulesEngine1";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.rulesEngines.get(resourceGroupName, frontDoorName, rulesEngineName);
  console.log(result);
}

getRulesEngineConfiguration().catch(console.error);
