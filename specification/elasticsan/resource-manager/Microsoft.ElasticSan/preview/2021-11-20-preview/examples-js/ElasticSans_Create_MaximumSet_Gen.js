const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create ElasticSan.
 *
 * @summary Create ElasticSan.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/ElasticSans_Create_MaximumSet_Gen.json
 */
async function elasticSansCreateMaximumSetGen() {
  const subscriptionId = "aaaaaaaaaaaaaaaaaa";
  const resourceGroupName = "rgelasticsan";
  const elasticSanName = "ti7q-k952-1qB3J_5";
  const parameters = {
    availabilityZones: ["aaaaaaaaaaaaaaaaa"],
    baseSizeTiB: 26,
    extendedCapacitySizeTiB: 7,
    location: "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
    sku: { name: "Premium_LRS", tier: "Premium" },
    tags: { key896: "aaaaaaaaaaaaaaaaaa" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.elasticSans.beginCreateAndWait(
    resourceGroupName,
    elasticSanName,
    parameters
  );
  console.log(result);
}

elasticSansCreateMaximumSetGen().catch(console.error);
