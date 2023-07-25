const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The configuration object for the specified cluster. This API is not recommended and might be removed in the future. Please consider using List configurations API instead.
 *
 * @summary The configuration object for the specified cluster. This API is not recommended and might be removed in the future. Please consider using List configurations API instead.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/HDI_Configurations_Get.json
 */
async function getCoreSiteSettings() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const configurationName = "core-site";
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.configurations.get(resourceGroupName, clusterName, configurationName);
  console.log(result);
}
