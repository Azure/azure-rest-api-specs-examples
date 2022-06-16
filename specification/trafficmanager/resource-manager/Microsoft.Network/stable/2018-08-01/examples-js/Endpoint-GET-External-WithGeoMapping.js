const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

async function endpointGetExternalWithGeoMapping() {
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

endpointGetExternalWithGeoMapping().catch(console.error);
