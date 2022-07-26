const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the status of operation.
 *
 * @summary Gets the status of operation.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/OperationsStatusGet.json
 */
async function operationsStatusGet() {
  const subscriptionId = "subid";
  const location = "West US";
  const operationId = "testoperationid";
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.operationsStatus.get(location, operationId);
  console.log(result);
}

operationsStatusGet().catch(console.error);
