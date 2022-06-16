const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the attached database configuration resource name is valid and is not already in use.
 *
 * @summary Checks that the attached database configuration resource name is valid and is not already in use.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoAttachedDatabaseConfigurationCheckNameAvailability.json
 */
async function kustoAttachedDatabaseConfigurationCheckNameAvailability() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const resourceName = {
    name: "adc1",
    type: "Microsoft.Kusto/clusters/attachedDatabaseConfigurations",
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.attachedDatabaseConfigurations.checkNameAvailability(
    resourceGroupName,
    clusterName,
    resourceName
  );
  console.log(result);
}

kustoAttachedDatabaseConfigurationCheckNameAvailability().catch(console.error);
