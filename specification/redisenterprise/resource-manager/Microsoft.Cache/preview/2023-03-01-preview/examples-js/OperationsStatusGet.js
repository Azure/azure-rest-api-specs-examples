const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the status of operation.
 *
 * @summary Gets the status of operation.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2023-03-01-preview/examples/OperationsStatusGet.json
 */
async function operationsStatusGet() {
  const subscriptionId = process.env["REDISENTERPRISE_SUBSCRIPTION_ID"] || "subid";
  const location = "West US";
  const operationId = "testoperationid";
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.operationsStatus.get(location, operationId);
  console.log(result);
}
