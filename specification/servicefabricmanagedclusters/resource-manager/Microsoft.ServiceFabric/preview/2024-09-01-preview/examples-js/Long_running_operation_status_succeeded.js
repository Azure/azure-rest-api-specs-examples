const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get long running operation status.
 *
 * @summary Get long running operation status.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/examples/Long_running_operation_status_succeeded.json
 */
async function getSucceededOperationResult() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const operationId = "00000000-0000-0000-0000-000000001234";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.operationStatus.get(location, operationId);
  console.log(result);
}
