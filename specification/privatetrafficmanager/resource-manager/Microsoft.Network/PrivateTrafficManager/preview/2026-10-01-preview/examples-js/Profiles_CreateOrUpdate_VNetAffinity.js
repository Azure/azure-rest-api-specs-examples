const { PrivateTrafficManagerManagementClient } = require("@azure/arm-privatetrafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Private Traffic Manager profile.
 *
 * @summary create or update a Private Traffic Manager profile.
 * x-ms-original-file: 2026-10-01-preview/Profiles_CreateOrUpdate_VNetAffinity.json
 */
async function profilesCreateOrUpdateVNetAffinity() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA";
  const client = new PrivateTrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.profiles.createOrUpdate(
    "rgprivateTrafficManager",
    "vnet-affinity-profile",
    {
      location: "global",
      properties: {
        customTopologyMapMode: "Disabled",
        profileStatus: "Enabled",
        trafficRoutingMethod: "VNet",
        endpoints: [
          {
            name: "matched-endpoint",
            target: "10.0.0.10",
            endpointStatus: "Enabled",
            virtualNetworkIds: [
              "/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/virtualNetworks/vnet-east",
            ],
          },
          { name: "fallback-endpoint", target: "10.0.0.20", endpointStatus: "Enabled" },
        ],
      },
    },
  );
  console.log(result);
}
