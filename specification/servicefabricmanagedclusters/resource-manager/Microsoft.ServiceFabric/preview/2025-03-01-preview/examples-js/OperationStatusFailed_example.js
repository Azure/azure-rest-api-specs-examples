const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get long running operation status.
 *
 * @summary get long running operation status.
 * x-ms-original-file: 2025-03-01-preview/OperationStatusFailed_example.json
 */
async function errorResponseDescribingWhyTheOperationFailed() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.operationStatus.get("eastus", "00000000-0000-0000-0000-000000001234");
  console.log(result);
}
