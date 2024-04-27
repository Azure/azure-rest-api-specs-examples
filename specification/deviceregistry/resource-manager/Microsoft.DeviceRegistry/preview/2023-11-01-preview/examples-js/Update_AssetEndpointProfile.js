const { DeviceRegistryManagementClient } = require("@azure/arm-deviceregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a AssetEndpointProfile
 *
 * @summary Update a AssetEndpointProfile
 * x-ms-original-file: specification/deviceregistry/resource-manager/Microsoft.DeviceRegistry/preview/2023-11-01-preview/examples/Update_AssetEndpointProfile.json
 */
async function patchAnAssetEndpointProfile() {
  const subscriptionId =
    process.env["DEVICEREGISTRY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["DEVICEREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const assetEndpointProfileName = "my-assetendpointprofile";
  const properties = {
    properties: { targetAddress: "https://www.example.com/myTargetAddress" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DeviceRegistryManagementClient(credential, subscriptionId);
  const result = await client.assetEndpointProfiles.beginUpdateAndWait(
    resourceGroupName,
    assetEndpointProfileName,
    properties,
  );
  console.log(result);
}
