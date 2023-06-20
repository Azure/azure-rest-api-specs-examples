const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Traffic Manager endpoint.
 *
 * @summary Create or update a Traffic Manager endpoint.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Endpoint-PUT-External-WithCustomHeaders.json
 */
async function endpointPutExternalWithCustomHeaders() {
  const subscriptionId = process.env["TRAFFICMANAGER_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName =
    process.env["TRAFFICMANAGER_RESOURCE_GROUP"] || "azuresdkfornetautoresttrafficmanager1421";
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
