const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

async function endpointPutExternalWithLocation() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "azuresdkfornetautoresttrafficmanager1421";
  const profileName = "azsmnet6386";
  const endpointType = "ExternalEndpoints";
  const endpointName = "azsmnet7187";
  const parameters = {
    name: "azsmnet7187",
    type: "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
    endpointLocation: "North Europe",
    endpointStatus: "Enabled",
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

endpointPutExternalWithLocation().catch(console.error);
