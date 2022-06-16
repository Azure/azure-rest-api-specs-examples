const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Traffic Manager profile.
 *
 * @summary Create or update a Traffic Manager profile.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Profile-PUT-MultiValue.json
 */
async function profilePutMultiValue() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "azuresdkfornetautoresttrafficmanager1421";
  const profileName = "azsmnet6386";
  const parameters = {
    dnsConfig: { relativeName: "azsmnet6386", ttl: 35 },
    location: "global",
    maxReturn: 2,
    monitorConfig: { path: "/testpath.aspx", port: 80, protocol: "HTTP" },
    profileStatus: "Enabled",
    trafficRoutingMethod: "MultiValue",
    trafficViewEnrollmentStatus: "Disabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.profiles.createOrUpdate(resourceGroupName, profileName, parameters);
  console.log(result);
}

profilePutMultiValue().catch(console.error);
