const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create ElasticSan.
 *
 * @summary create ElasticSan.
 * x-ms-original-file: 2026-05-01-preview/ElasticSans_V2_Create_MaximumSet_Gen.json
 */
async function elasticSansV2CreateMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.elasticSans.create("resourcegroupname", "elasticsanname", {
    location: "France Central",
    properties: {
      autoScaleProperties: {
        scaleUpProperties: {
          autoScalePolicyEnforcement: "None",
          capacityUnitScaleUpLimitTiB: 17,
          increaseCapacityUnitByTiB: 4,
          unusedSizeTiB: 24,
        },
      },
      availabilityZones: ["1"],
      version: "V2",
      totalIops: 22,
      totalMBps: 4,
      totalSizeTiB: 27,
      publicNetworkAccess: "Enabled",
      sku: { name: "ElasticSAN_LRS" },
    },
    tags: { key9316: "ihndtieqibtob" },
  });
  console.log(result);
}
