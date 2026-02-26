const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the specified private endpoint connection associated with the Elastic San
 *
 * @summary gets the specified private endpoint connection associated with the Elastic San
 * x-ms-original-file: 2025-09-01/PrivateEndpointConnections_Get_MaximumSet_Gen.json
 */
async function privateEndpointConnectionsGetMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.privateEndpointConnections.get(
    "resourcegroupname",
    "elasticsanname",
    "privateendpointconnectionname",
  );
  console.log(result);
}
