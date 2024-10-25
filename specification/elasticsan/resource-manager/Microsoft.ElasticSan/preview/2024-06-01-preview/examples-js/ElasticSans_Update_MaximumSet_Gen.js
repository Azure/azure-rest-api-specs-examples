const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a Elastic San.
 *
 * @summary Update a Elastic San.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/ElasticSans_Update_MaximumSet_Gen.json
 */
async function elasticSansUpdateMaximumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const parameters = {
    properties: {
      autoScaleProperties: {
        scaleUpProperties: {
          autoScalePolicyEnforcement: "None",
          capacityUnitScaleUpLimitTiB: 17,
          increaseCapacityUnitByTiB: 4,
          unusedSizeTiB: 24,
        },
      },
      baseSizeTiB: 21,
      extendedCapacitySizeTiB: 10,
      publicNetworkAccess: "Enabled",
    },
    tags: { key5945: "eufqmxresekerilkgvfzddolxzas" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.elasticSans.beginUpdateAndWait(
    resourceGroupName,
    elasticSanName,
    parameters,
  );
  console.log(result);
}
