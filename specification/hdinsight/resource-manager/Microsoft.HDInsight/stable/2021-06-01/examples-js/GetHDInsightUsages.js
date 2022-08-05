const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the usages for the specified location.
 *
 * @summary Lists the usages for the specified location.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetHDInsightUsages.json
 */
async function getTheSubscriptionUsagesForSpecificLocation() {
  const subscriptionId = "subid";
  const location = "West US";
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.locations.listUsages(location);
  console.log(result);
}

getTheSubscriptionUsagesForSpecificLocation().catch(console.error);
