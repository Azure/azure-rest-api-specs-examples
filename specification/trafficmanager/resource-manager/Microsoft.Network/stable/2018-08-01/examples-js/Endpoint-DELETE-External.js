const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Traffic Manager endpoint.
 *
 * @summary Deletes a Traffic Manager endpoint.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Endpoint-DELETE-External.json
 */
async function endpointDeleteExternal() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "azuresdkfornetautoresttrafficmanager1421";
  const profileName = "azsmnet6386";
  const endpointType = "ExternalEndpoints";
  const endpointName = "azsmnet7187";
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.endpoints.delete(
    resourceGroupName,
    profileName,
    endpointType,
    endpointName
  );
  console.log(result);
}

endpointDeleteExternal().catch(console.error);
