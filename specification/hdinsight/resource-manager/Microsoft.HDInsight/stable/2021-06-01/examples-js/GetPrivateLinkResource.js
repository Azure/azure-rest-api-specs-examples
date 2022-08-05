const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specific private link resource.
 *
 * @summary Gets the specific private link resource.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetPrivateLinkResource.json
 */
async function getSpecificPrivateLinkResourceInASpecificHdInsightCluster() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cluster1";
  const privateLinkResourceName = "gateway";
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.get(
    resourceGroupName,
    clusterName,
    privateLinkResourceName
  );
  console.log(result);
}

getSpecificPrivateLinkResourceInASpecificHdInsightCluster().catch(console.error);
