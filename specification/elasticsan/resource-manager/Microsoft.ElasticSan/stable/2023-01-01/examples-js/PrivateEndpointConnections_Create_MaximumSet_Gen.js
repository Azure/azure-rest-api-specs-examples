const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the state of specified private endpoint connection associated with the Elastic San
 *
 * @summary Update the state of specified private endpoint connection associated with the Elastic San
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/PrivateEndpointConnections_Create_MaximumSet_Gen.json
 */
async function privateEndpointConnectionsCreateMaximumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const privateEndpointConnectionName = "privateendpointconnectionname";
  const parameters = {
    properties: {
      groupIds: ["jdwrzpemdjrpiwzvy"],
      privateEndpoint: {},
      privateLinkServiceConnectionState: {
        description: "dxl",
        actionsRequired: "jhjdpwvyzipggtn",
        status: "Pending",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginCreateAndWait(
    resourceGroupName,
    elasticSanName,
    privateEndpointConnectionName,
    parameters,
  );
  console.log(result);
}
