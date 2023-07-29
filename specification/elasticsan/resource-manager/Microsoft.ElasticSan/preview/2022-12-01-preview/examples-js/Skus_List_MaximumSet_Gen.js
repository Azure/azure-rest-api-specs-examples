const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the available Skus in the region and information related to them
 *
 * @summary List all the available Skus in the region and information related to them
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2022-12-01-preview/examples/Skus_List_MaximumSet_Gen.json
 */
async function skusListMaximumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const filter = "dtycml";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.skus.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
