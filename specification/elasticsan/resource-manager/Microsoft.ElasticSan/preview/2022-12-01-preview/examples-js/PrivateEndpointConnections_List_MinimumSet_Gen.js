const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all Private Endpoint Connections associated with the Elastic San.
 *
 * @summary List all Private Endpoint Connections associated with the Elastic San.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2022-12-01-preview/examples/PrivateEndpointConnections_List_MinimumSet_Gen.json
 */
async function privateEndpointConnectionsListMinimumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.list(
    resourceGroupName,
    elasticSanName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
