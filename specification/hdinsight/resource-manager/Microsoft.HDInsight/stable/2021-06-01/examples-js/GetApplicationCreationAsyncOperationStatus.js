const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the async operation status.
 *
 * @summary Gets the async operation status.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetApplicationCreationAsyncOperationStatus.json
 */
async function getTheAzureAsyncOperationStatus() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cluster1";
  const applicationName = "app";
  const operationId = "CF938302-6B4D-44A0-A6D2-C0D67E847AEC";
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.applications.getAzureAsyncOperationStatus(
    resourceGroupName,
    clusterName,
    applicationName,
    operationId
  );
  console.log(result);
}

getTheAzureAsyncOperationStatus().catch(console.error);
