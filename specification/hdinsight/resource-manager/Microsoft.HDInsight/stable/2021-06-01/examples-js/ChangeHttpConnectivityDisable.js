const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Configures the HTTP settings on the specified cluster. This API is deprecated, please use UpdateGatewaySettings in cluster endpoint instead.
 *
 * @summary Configures the HTTP settings on the specified cluster. This API is deprecated, please use UpdateGatewaySettings in cluster endpoint instead.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/ChangeHttpConnectivityDisable.json
 */
async function disableHttpConnectivity() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cluster1";
  const configurationName = "gateway";
  const parameters = {
    restAuthCredentialIsEnabled: "false",
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.configurations.beginUpdateAndWait(
    resourceGroupName,
    clusterName,
    configurationName,
    parameters
  );
  console.log(result);
}

disableHttpConnectivity().catch(console.error);
