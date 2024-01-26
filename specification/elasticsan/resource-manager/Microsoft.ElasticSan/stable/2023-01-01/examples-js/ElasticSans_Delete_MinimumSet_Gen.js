const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a Elastic San.
 *
 * @summary Delete a Elastic San.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/ElasticSans_Delete_MinimumSet_Gen.json
 */
async function elasticSansDeleteMinimumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.elasticSans.beginDeleteAndWait(resourceGroupName, elasticSanName);
  console.log(result);
}
