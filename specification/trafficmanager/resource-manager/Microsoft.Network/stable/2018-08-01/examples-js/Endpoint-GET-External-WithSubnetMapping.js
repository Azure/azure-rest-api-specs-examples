const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Traffic Manager endpoint.
 *
 * @summary Gets a Traffic Manager endpoint.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Endpoint-GET-External-WithSubnetMapping.json
 */
async function endpointGetExternalWithSubnetMapping() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "azuresdkfornetautoresttrafficmanager2191";
  const profileName = "azuresdkfornetautoresttrafficmanager8224";
  const endpointType = "ExternalEndpoints";
  const endpointName = "My%20external%20endpoint";
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.endpoints.get(
    resourceGroupName,
    profileName,
    endpointType,
    endpointName
  );
  console.log(result);
}

endpointGetExternalWithSubnetMapping().catch(console.error);
