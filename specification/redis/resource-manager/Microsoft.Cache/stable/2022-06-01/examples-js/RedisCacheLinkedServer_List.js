const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of linked servers associated with this redis cache (requires Premium SKU).
 *
 * @summary Gets the list of linked servers associated with this redis cache (requires Premium SKU).
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheLinkedServer_List.json
 */
async function linkedServerList() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const name = "cache1";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.linkedServer.list(resourceGroupName, name)) {
    resArray.push(item);
  }
  console.log(resArray);
}

linkedServerList().catch(console.error);
