const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the capabilities for the specified location.
 *
 * @summary Gets the capabilities for the specified location.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetHDInsightCapabilities.json
 */
async function getTheSubscriptionCapabilitiesForSpecificLocation() {
  const subscriptionId = "subid";
  const location = "West US";
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.locations.getCapabilities(location);
  console.log(result);
}

getTheSubscriptionCapabilitiesForSpecificLocation().catch(console.error);
