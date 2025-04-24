const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Delete the provided cloud services network.
 *
 * @summary Delete the provided cloud services network.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-10-01-preview/examples/CloudServicesNetworks_Delete.json
 */
async function deleteCloudServicesNetwork() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const cloudServicesNetworkName = "cloudServicesNetworkName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.cloudServicesNetworks.beginDeleteAndWait(
    resourceGroupName,
    cloudServicesNetworkName,
  );
  console.log(result);
}
