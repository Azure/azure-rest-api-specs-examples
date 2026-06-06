const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a Traffic Manager endpoint.
 *
 * @summary gets a Traffic Manager endpoint.
 * x-ms-original-file: 2024-04-01-preview/Endpoint-GET-External-WithGeoMapping.json
 */
async function endpointGETExternalWithGeoMapping() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.endpoints.get(
    "azuresdkfornetautoresttrafficmanager2191",
    "azuresdkfornetautoresttrafficmanager8224",
    "ExternalEndpoints",
    "My%20external%20endpoint",
  );
  console.log(result);
}
