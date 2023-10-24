const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to
 *
 * @summary
 * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/OperationsDiscovery_Get.json
 */
async function operationsDiscoveryGet() {
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential);
  const result = await client.operationsDiscoveryOperations.get();
  console.log(result);
}
