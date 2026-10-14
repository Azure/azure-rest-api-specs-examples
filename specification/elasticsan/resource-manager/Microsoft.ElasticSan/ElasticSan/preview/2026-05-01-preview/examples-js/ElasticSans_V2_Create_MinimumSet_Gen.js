const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create ElasticSan.
 *
 * @summary create ElasticSan.
 * x-ms-original-file: 2026-05-01-preview/ElasticSans_V2_Create_MinimumSet_Gen.json
 */
async function elasticSansV2CreateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.elasticSans.create("resourcegroupname", "elasticsanname", {
    location: "France Central",
    properties: {
      version: "V2",
      totalIops: 22,
      totalMBps: 4,
      totalSizeTiB: 27,
      sku: { name: "ElasticSAN_LRS" },
    },
  });
  console.log(result);
}
