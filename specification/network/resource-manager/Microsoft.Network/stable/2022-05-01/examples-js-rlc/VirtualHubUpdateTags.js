const createNetworkManagementClient = require("@azure-rest/arm-network").default;
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates VirtualHub tags.
 *
 * @summary Updates VirtualHub tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualHubUpdateTags.json
 */
async function virtualHubUpdate() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const virtualHubName = "virtualHub2";
  const options = {
    body: { tags: { key1: "value1", key2: "value2" } },
    queryParameters: { "api-version": "2022-05-01" },
  };
  const result = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}",
      subscriptionId,
      resourceGroupName,
      virtualHubName,
    )
    .patch(options);
  console.log(result);
}

virtualHubUpdate().catch(console.error);
