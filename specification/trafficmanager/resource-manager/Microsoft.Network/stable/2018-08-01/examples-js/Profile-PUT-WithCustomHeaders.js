const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Traffic Manager profile.
 *
 * @summary Create or update a Traffic Manager profile.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Profile-PUT-WithCustomHeaders.json
 */
async function profilePutWithCustomHeaders() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "azuresdkfornetautoresttrafficmanager2583";
  const profileName = "azuresdkfornetautoresttrafficmanager6192";
  const parameters = {
    dnsConfig: {
      relativeName: "azuresdkfornetautoresttrafficmanager6192",
      ttl: 35,
    },
    endpoints: [
      {
        name: "My external endpoint",
        type: "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
        customHeaders: [{ name: "header-2", value: "value-2-overridden" }],
        endpointLocation: "North Europe",
        endpointStatus: "Enabled",
        target: "foobar.contoso.com",
      },
    ],
    location: "global",
    monitorConfig: {
      path: "/testpath.aspx",
      customHeaders: [
        { name: "header-1", value: "value-1" },
        { name: "header-2", value: "value-2" },
      ],
      expectedStatusCodeRanges: [
        { max: 205, min: 200 },
        { max: 410, min: 400 },
      ],
      intervalInSeconds: 10,
      port: 80,
      timeoutInSeconds: 5,
      toleratedNumberOfFailures: 2,
      protocol: "HTTP",
    },
    profileStatus: "Enabled",
    trafficRoutingMethod: "Performance",
    trafficViewEnrollmentStatus: "Disabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.profiles.createOrUpdate(resourceGroupName, profileName, parameters);
  console.log(result);
}

profilePutWithCustomHeaders().catch(console.error);
