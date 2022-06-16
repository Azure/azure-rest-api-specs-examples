const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Frontend endpoint with the specified name within the specified Front Door.
 *
 * @summary Gets a Frontend endpoint with the specified name within the specified Front Door.
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorFrontendEndpointGet.json
 */
async function getFrontendEndpoint() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const frontDoorName = "frontDoor1";
  const frontendEndpointName = "frontendEndpoint1";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.frontendEndpoints.get(
    resourceGroupName,
    frontDoorName,
    frontendEndpointName
  );
  console.log(result);
}

getFrontendEndpoint().catch(console.error);
