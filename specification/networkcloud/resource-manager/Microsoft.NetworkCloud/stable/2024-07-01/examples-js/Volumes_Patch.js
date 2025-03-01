const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update tags associated with the provided volume.
 *
 * @summary Update tags associated with the provided volume.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/Volumes_Patch.json
 */
async function patchVolume() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const volumeName = "volumeName";
  const volumeUpdateParameters = {
    tags: { key1: "myvalue1", key2: "myvalue2" },
  };
  const options = { volumeUpdateParameters };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.volumes.update(resourceGroupName, volumeName, options);
  console.log(result);
}
