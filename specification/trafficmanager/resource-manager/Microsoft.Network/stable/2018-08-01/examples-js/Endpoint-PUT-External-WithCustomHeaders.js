const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

async function endpointPutExternalWithCustomHeaders() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "azuresdkfornetautoresttrafficmanager1421";
  const profileName = "azsmnet6386";
  const endpointType = "ExternalEndpoints";
  const endpointName = "azsmnet7187";
  const parameters = {
    name: "azsmnet7187",
    type: "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
    customHeaders: [
      { name: "header-1", value: "value-1" },
      { name: "header-2", value: "value-2" },
    ],
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

endpointPutExternalWithCustomHeaders().catch(console.error);
