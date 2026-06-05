const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Traffic Manager endpoint.
 *
 * @summary create or update a Traffic Manager endpoint.
 * x-ms-original-file: 2024-04-01-preview/Endpoint-PUT-External-WithSubnetMapping.json
 */
async function endpointPUTExternalWithSubnetMapping() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.endpoints.createOrUpdate(
    "azuresdkfornetautoresttrafficmanager2191",
    "azuresdkfornetautoresttrafficmanager8224",
    "ExternalEndpoints",
    "My%20external%20endpoint",
    {
      name: "My external endpoint",
      type: "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
      endpointStatus: "Enabled",
      subnets: [
        { first: "1.2.3.0", scope: 24 },
        { first: "25.26.27.28", last: "29.30.31.32" },
      ],
      target: "foobar.contoso.com",
    },
  );
  console.log(result);
}
