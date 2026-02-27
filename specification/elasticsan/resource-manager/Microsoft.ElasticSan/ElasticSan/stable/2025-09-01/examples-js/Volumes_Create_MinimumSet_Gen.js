const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a Volume.
 *
 * @summary create a Volume.
 * x-ms-original-file: 2025-09-01/Volumes_Create_MinimumSet_Gen.json
 */
async function volumesCreateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.volumes.create(
    "resourcegroupname",
    "elasticsanname",
    "volumegroupname",
    "volumename",
    { properties: { sizeGiB: 9 } },
  );
  console.log(result);
}
