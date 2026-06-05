const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Traffic Manager endpoint.
 *
 * @summary create or update a Traffic Manager endpoint.
 * x-ms-original-file: 2024-04-01-preview/Endpoint-PUT-External-WithAlwaysServe.json
 */
async function endpointPUTExternalWithAlwaysServe() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.endpoints.createOrUpdate(
    "azuresdkfornetautoresttrafficmanager1421",
    "azsmnet6386",
    "ExternalEndpoints",
    "azsmnet7187",
    {
      name: "azsmnet7187",
      type: "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
      alwaysServe: "Enabled",
      endpointLocation: "North Europe",
      endpointStatus: "Enabled",
      target: "foobar.contoso.com",
    },
  );
  console.log(result);
}
