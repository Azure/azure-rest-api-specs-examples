const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a managed private endpoint.
 *
 * @summary Updates a managed private endpoint.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoManagedPrivateEndpointsUpdate.json
 */
async function kustoManagedPrivateEndpointsUpdate() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const managedPrivateEndpointName = "managedPrivateEndpointTest";
  const parameters = {
    groupId: "blob",
    privateLinkResourceId:
      "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/storageAccountTest",
    requestMessage: "Please Approve Managed Private Endpoint Request.",
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.managedPrivateEndpoints.beginUpdateAndWait(
    resourceGroupName,
    clusterName,
    managedPrivateEndpointName,
    parameters
  );
  console.log(result);
}

kustoManagedPrivateEndpointsUpdate().catch(console.error);
