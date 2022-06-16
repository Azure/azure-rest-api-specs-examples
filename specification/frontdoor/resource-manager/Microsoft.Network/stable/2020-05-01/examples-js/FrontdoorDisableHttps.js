const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Disables a frontendEndpoint for HTTPS traffic
 *
 * @summary Disables a frontendEndpoint for HTTPS traffic
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorDisableHttps.json
 */
async function frontendEndpointsDisableHttps() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const frontDoorName = "frontDoor1";
  const frontendEndpointName = "frontendEndpoint1";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.frontendEndpoints.beginDisableHttpsAndWait(
    resourceGroupName,
    frontDoorName,
    frontendEndpointName
  );
  console.log(result);
}

frontendEndpointsDisableHttps().catch(console.error);
