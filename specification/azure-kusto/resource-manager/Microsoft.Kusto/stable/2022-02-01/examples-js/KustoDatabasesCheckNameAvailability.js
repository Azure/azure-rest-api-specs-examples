const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the databases resource name is valid and is not already in use.
 *
 * @summary Checks that the databases resource name is valid and is not already in use.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDatabasesCheckNameAvailability.json
 */
async function kustoDatabasesCheckNameAvailability() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const resourceName = {
    name: "database1",
    type: "Microsoft.Kusto/clusters/databases",
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.databases.checkNameAvailability(
    resourceGroupName,
    clusterName,
    resourceName
  );
  console.log(result);
}

kustoDatabasesCheckNameAvailability().catch(console.error);
