const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a Elastic San.
 *
 * @summary update a Elastic San.
 * x-ms-original-file: 2026-05-01-preview/ElasticSans_V1_Update_MinimumSet_Gen.json
 */
async function elasticSansV1UpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.elasticSans.update("resourcegroupname", "elasticsanname", {});
  console.log(result);
}
