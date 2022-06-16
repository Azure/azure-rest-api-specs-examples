const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

async function endpointPutExternalWithGeoMapping() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "azuresdkfornetautoresttrafficmanager2191";
  const profileName = "azuresdkfornetautoresttrafficmanager8224";
  const endpointType = "ExternalEndpoints";
  const endpointName = "My%20external%20endpoint";
  const parameters = {
    name: "My external endpoint",
    type: "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
    endpointStatus: "Enabled",
    geoMapping: ["GEO-AS", "GEO-AF"],
    target: "foobar.contoso.com",
  };
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.endpoints.createOrUpdate(
    resourceGroupName,
    profileName,
    endpointType,
    endpointName,
    parameters
  );
  console.log(result);
}

endpointPutExternalWithGeoMapping().catch(console.error);
