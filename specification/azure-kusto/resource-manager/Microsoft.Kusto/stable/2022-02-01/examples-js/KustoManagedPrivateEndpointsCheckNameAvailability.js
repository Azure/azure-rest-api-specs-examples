const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the managed private endpoints resource name is valid and is not already in use.
 *
 * @summary Checks that the managed private endpoints resource name is valid and is not already in use.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoManagedPrivateEndpointsCheckNameAvailability.json
 */
async function kustoManagedPrivateEndpointsCheckNameAvailability() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const resourceName = {
    name: "pme1",
    type: "Microsoft.Kusto/clusters/managedPrivateEndpoints",
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.managedPrivateEndpoints.checkNameAvailability(
    resourceGroupName,
    clusterName,
    resourceName
  );
  console.log(result);
}

kustoManagedPrivateEndpointsCheckNameAvailability().catch(console.error);
