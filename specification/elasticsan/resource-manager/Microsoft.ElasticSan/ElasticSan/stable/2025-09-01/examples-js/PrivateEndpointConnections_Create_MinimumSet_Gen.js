const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update the state of specified private endpoint connection associated with the Elastic San
 *
 * @summary update the state of specified private endpoint connection associated with the Elastic San
 * x-ms-original-file: 2025-09-01/PrivateEndpointConnections_Create_MinimumSet_Gen.json
 */
async function privateEndpointConnectionsCreateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.privateEndpointConnections.create(
    "resourcegroupname",
    "elasticsanname",
    "privateendpointconnectionname",
    { properties: { privateLinkServiceConnectionState: {} } },
  );
  console.log(result);
}
