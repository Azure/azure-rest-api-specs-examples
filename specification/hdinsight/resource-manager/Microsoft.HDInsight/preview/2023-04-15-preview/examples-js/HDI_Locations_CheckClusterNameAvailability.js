const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the cluster name is available or not.
 *
 * @summary Check the cluster name is available or not.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/HDI_Locations_CheckClusterNameAvailability.json
 */
async function getTheSubscriptionUsagesForSpecificLocation() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const location = "westus";
  const parameters = {
    name: "test123",
    type: "clusters",
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.locations.checkNameAvailability(location, parameters);
  console.log(result);
}
