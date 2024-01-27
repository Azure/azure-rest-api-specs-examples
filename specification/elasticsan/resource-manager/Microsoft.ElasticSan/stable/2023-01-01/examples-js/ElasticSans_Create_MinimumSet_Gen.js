const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create ElasticSan.
 *
 * @summary Create ElasticSan.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/ElasticSans_Create_MinimumSet_Gen.json
 */
async function elasticSansCreateMinimumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const parameters = {
    location: "France Central",
    properties: {
      baseSizeTiB: 15,
      extendedCapacitySizeTiB: 27,
      sku: { name: "Premium_LRS" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.elasticSans.beginCreateAndWait(
    resourceGroupName,
    elasticSanName,
    parameters,
  );
  console.log(result);
}
