const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a managed private endpoint.
 *
 * @summary Deletes a managed private endpoint.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoManagedPrivateEndpointsDelete.json
 */
async function managedPrivateEndpointsDelete() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const managedPrivateEndpointName = "managedPrivateEndpointTest";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.managedPrivateEndpoints.beginDeleteAndWait(
    resourceGroupName,
    clusterName,
    managedPrivateEndpointName
  );
  console.log(result);
}

managedPrivateEndpointsDelete().catch(console.error);
