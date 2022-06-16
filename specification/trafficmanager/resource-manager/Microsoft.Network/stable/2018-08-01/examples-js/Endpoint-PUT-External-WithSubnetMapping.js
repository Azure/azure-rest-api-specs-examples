const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

async function endpointPutExternalWithSubnetMapping() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "azuresdkfornetautoresttrafficmanager2191";
  const profileName = "azuresdkfornetautoresttrafficmanager8224";
  const endpointType = "ExternalEndpoints";
  const endpointName = "My%20external%20endpoint";
  const parameters = {
    name: "My external endpoint",
    type: "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
    endpointStatus: "Enabled",
    subnets: [
      { first: "1.2.3.0", scope: 24 },
      { first: "25.26.27.28", last: "29.30.31.32" },
    ],
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

endpointPutExternalWithSubnetMapping().catch(console.error);
