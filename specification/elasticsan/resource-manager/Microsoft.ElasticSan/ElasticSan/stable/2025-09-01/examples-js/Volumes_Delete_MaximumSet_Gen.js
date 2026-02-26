const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete an Volume.
 *
 * @summary delete an Volume.
 * x-ms-original-file: 2025-09-01/Volumes_Delete_MaximumSet_Gen.json
 */
async function volumesDeleteMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  await client.volumes.delete(
    "resourcegroupname",
    "elasticsanname",
    "volumegroupname",
    "volumename",
    { xMsDeleteSnapshots: "true", xMsForceDelete: "true" },
  );
}
