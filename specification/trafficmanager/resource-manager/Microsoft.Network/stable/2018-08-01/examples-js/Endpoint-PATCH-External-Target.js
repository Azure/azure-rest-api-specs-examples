const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

async function endpointPatchExternalTarget() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "azuresdkfornetautoresttrafficmanager1421";
  const profileName = "azsmnet6386";
  const endpointType = "ExternalEndpoints";
  const endpointName = "azsmnet7187";
  const parameters = {
    name: "azsmnet7187",
    type: "Microsoft.Network/trafficManagerProfiles/externalEndpoints",
    id: "/subscriptions/{subscription-id}/resourceGroups/azuresdkfornetautoresttrafficmanager1421/providers/Microsoft.Network/trafficManagerProfiles/azsmnet6386/externalEndpoints/azsmnet7187",
    target: "another.foobar.contoso.com",
  };
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.endpoints.update(
    resourceGroupName,
    profileName,
    endpointType,
    endpointName,
    parameters
  );
  console.log(result);
}

endpointPatchExternalTarget().catch(console.error);
