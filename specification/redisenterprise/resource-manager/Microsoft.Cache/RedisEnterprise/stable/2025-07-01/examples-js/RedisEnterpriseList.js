const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all Redis Enterprise clusters in the specified subscription.
 *
 * @summary Lists all Redis Enterprise clusters in the specified subscription.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/RedisEnterprise/stable/2025-07-01/examples/RedisEnterpriseList.json
 */
async function redisEnterpriseList() {
  const subscriptionId =
    process.env["REDISENTERPRISE_SUBSCRIPTION_ID"] || "e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f";
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.redisEnterprise.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
