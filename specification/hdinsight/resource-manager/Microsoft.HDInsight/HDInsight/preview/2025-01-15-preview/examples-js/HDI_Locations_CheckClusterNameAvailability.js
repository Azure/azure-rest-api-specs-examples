const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to check the cluster name is available or not.
 *
 * @summary check the cluster name is available or not.
 * x-ms-original-file: 2025-01-15-preview/HDI_Locations_CheckClusterNameAvailability.json
 */
async function getTheSubscriptionUsagesForSpecificLocation() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.locations.checkNameAvailability("westus", {
    name: "test123",
    type: "clusters",
  });
  console.log(result);
}
