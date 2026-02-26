const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a Elastic San.
 *
 * @summary delete a Elastic San.
 * x-ms-original-file: 2025-09-01/ElasticSans_Delete_MaximumSet_Gen.json
 */
async function elasticSansDeleteMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  await client.elasticSans.delete("resourcegroupname", "elasticsanname");
}
