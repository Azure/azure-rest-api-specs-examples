const { DeviceRegistryManagementClient } = require("@azure/arm-deviceregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a AssetEndpointProfile
 *
 * @summary Create a AssetEndpointProfile
 * x-ms-original-file: specification/deviceregistry/resource-manager/Microsoft.DeviceRegistry/preview/2023-11-01-preview/examples/Create_AssetEndpointProfile.json
 */
async function createAnAssetEndpointProfile() {
  const subscriptionId =
    process.env["DEVICEREGISTRY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["DEVICEREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const assetEndpointProfileName = "my-assetendpointprofile";
  const resource = {
    extendedLocation: {
      name: "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1",
      type: "CustomLocation",
    },
    location: "West Europe",
    properties: {
      targetAddress: "https://www.example.com/myTargetAddress",
      userAuthentication: { mode: "Anonymous" },
    },
    tags: { site: "building-1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DeviceRegistryManagementClient(credential, subscriptionId);
  const result = await client.assetEndpointProfiles.beginCreateOrReplaceAndWait(
    resourceGroupName,
    assetEndpointProfileName,
    resource,
  );
  console.log(result);
}
