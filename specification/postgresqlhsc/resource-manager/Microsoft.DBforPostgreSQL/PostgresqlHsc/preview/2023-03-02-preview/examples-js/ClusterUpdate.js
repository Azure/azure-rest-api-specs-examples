const { DBforPostgreSQLClient } = require("@azure/arm-postgresqlhsc");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing cluster. The request body can contain one or several properties from the cluster definition.
 *
 * @summary updates an existing cluster. The request body can contain one or several properties from the cluster definition.
 * x-ms-original-file: 2023-03-02-preview/ClusterUpdate.json
 */
async function updateMultipleConfigurationSettingsOfTheCluster() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new DBforPostgreSQLClient(credential, subscriptionId);
  const result = await client.clusters.update("TestGroup", "testcluster", {
    administratorLoginPassword: "newpassword",
    coordinatorVCores: 16,
    nodeCount: 4,
    nodeVCores: 16,
  });
  console.log(result);
}
