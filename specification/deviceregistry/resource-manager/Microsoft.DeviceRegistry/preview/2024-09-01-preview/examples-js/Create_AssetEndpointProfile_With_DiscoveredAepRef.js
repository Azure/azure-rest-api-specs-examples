const { DeviceRegistryManagementClient } = require("@azure/arm-deviceregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a AssetEndpointProfile
 *
 * @summary create a AssetEndpointProfile
 * x-ms-original-file: 2024-09-01-preview/Create_AssetEndpointProfile_With_DiscoveredAepRef.json
 */
async function createAssetEndpointProfileWithDiscoveredAepRef() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DeviceRegistryManagementClient(credential, subscriptionId);
  const result = await client.assetEndpointProfiles.createOrReplace(
    "myResourceGroup",
    "my-assetendpointprofile",
    {
      location: "West Europe",
      extendedLocation: {
        type: "CustomLocation",
        name: "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1",
      },
      tags: { site: "building-1" },
      properties: {
        targetAddress: "https://www.example.com/myTargetAddress",
        endpointProfileType: "myEndpointProfileType",
        discoveredAssetEndpointProfileRef: "discoveredAssetEndpointProfile1",
        authentication: { method: "Anonymous" },
      },
    },
  );
  console.log(result);
}
