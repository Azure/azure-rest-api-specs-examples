const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update tags associated with the provided trunked network.
 *
 * @summary Update tags associated with the provided trunked network.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/TrunkedNetworks_Patch.json
 */
async function patchTrunkedNetwork() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const trunkedNetworkName = "trunkedNetworkName";
  const trunkedNetworkUpdateParameters = {
    tags: { key1: "myvalue1", key2: "myvalue2" },
  };
  const options = {
    trunkedNetworkUpdateParameters,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.trunkedNetworks.update(
    resourceGroupName,
    trunkedNetworkName,
    options,
  );
  console.log(result);
}
