const { DeviceRegistryManagementClient } = require("@azure/arm-deviceregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a AssetEndpointProfile
 *
 * @summary update a AssetEndpointProfile
 * x-ms-original-file: 2024-09-01-preview/Update_AssetEndpointProfile.json
 */
async function updateAssetEndpointProfile() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DeviceRegistryManagementClient(credential, subscriptionId);
  const result = await client.assetEndpointProfiles.update(
    "myResourceGroup",
    "my-assetendpointprofile",
    {
      properties: { targetAddress: "https://www.example.com/myTargetAddress" },
    },
  );
  console.log(result);
}
