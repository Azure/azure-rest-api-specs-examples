const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all configuration information for an HDI cluster.
 *
 * @summary Gets all configuration information for an HDI cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/HDI_Configurations_List.json
 */
async function getAllConfigurationInformation() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cluster1";
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.configurations.list(resourceGroupName, clusterName);
  console.log(result);
}

getAllConfigurationInformation().catch(console.error);
