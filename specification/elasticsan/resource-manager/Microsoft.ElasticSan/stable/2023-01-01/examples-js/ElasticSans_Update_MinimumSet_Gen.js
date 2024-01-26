const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a Elastic San.
 *
 * @summary Update a Elastic San.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/ElasticSans_Update_MinimumSet_Gen.json
 */
async function elasticSansUpdateMinimumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const parameters = {};
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.elasticSans.beginUpdateAndWait(
    resourceGroupName,
    elasticSanName,
    parameters,
  );
  console.log(result);
}
