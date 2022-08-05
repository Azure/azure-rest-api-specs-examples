const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specific private endpoint connection.
 *
 * @summary Gets the specific private endpoint connection.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetPrivateEndpointConnection.json
 */
async function getSpecificPrivateEndpointConnectionForASpecificHdInsightCluster() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cluster1";
  const privateEndpointConnectionName = "testprivateep.b3bf5fed-9b12-4560-b7d0-2abe1bba07e2";
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.get(
    resourceGroupName,
    clusterName,
    privateEndpointConnectionName
  );
  console.log(result);
}

getSpecificPrivateEndpointConnectionForASpecificHdInsightCluster().catch(console.error);
