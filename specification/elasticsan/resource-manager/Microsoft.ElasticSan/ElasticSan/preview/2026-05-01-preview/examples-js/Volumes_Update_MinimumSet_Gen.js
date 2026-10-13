const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update an Volume.
 *
 * @summary update an Volume.
 * x-ms-original-file: 2026-05-01-preview/Volumes_Update_MinimumSet_Gen.json
 */
async function volumesUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.volumes.update(
    "resourcegroupname",
    "elasticsanname",
    "volumegroupname",
    "volumename",
    {},
  );
  console.log(result);
}
