const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the async operation status.
 *
 * @summary Get the async operation status.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/HDI_Locations_GetAsyncOperationStatus.json
 */
async function getsTheAzureAsyncOperationStatus() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const location = "East US 2";
  const operationId = "8a0348f4-8a85-4ec2-abe0-03b26104a9a0-0";
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.locations.getAzureAsyncOperationStatus(location, operationId);
  console.log(result);
}
