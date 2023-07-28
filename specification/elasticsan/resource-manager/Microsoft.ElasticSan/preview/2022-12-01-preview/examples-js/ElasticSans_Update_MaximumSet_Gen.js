const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a Elastic San.
 *
 * @summary Update a Elastic San.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2022-12-01-preview/examples/ElasticSans_Update_MaximumSet_Gen.json
 */
async function elasticSansUpdateMaximumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const parameters = {
    baseSizeTiB: 10,
    extendedCapacitySizeTiB: 22,
    tags: { key4212: "cqvcnwfefljntgeio" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.elasticSans.beginUpdateAndWait(
    resourceGroupName,
    elasticSanName,
    parameters
  );
  console.log(result);
}
