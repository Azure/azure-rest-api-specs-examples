const { ServiceFabricMeshManagementClient } = require("@azure/arm-servicefabricmesh");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a gateway resource with the specified name, description and properties. If a gateway resource with the same name exists, then it is updated with the specified description and properties. Use gateway resources to create a gateway for public connectivity for services within your application.
 *
 * @summary Creates a gateway resource with the specified name, description and properties. If a gateway resource with the same name exists, then it is updated with the specified description and properties. Use gateway resources to create a gateway for public connectivity for services within your application.
 * x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/gateways/create_update.json
 */
async function createOrUpdateGateway() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sbz_demo";
  const gatewayResourceName = "sampleGateway";
  const gatewayResourceDescription = {
    description: "Service Fabric Mesh sample gateway.",
    destinationNetwork: { name: "helloWorldNetwork" },
    location: "EastUS",
    sourceNetwork: { name: "Open" },
    tags: {},
    tcp: [
      {
        name: "web",
        destination: {
          applicationName: "helloWorldApp",
          endpointName: "helloWorldListener",
          serviceName: "helloWorldService",
        },
        port: 80,
      },
    ],
    http: [
      {
        name: "contosoWebsite",
        hosts: [
          {
            name: "contoso.com",
            routes: [
              {
                name: "index",
                destination: {
                  applicationName: "httpHelloWorldApp",
                  endpointName: "indexHttpEndpoint",
                  serviceName: "indexService",
                },
                match: {
                  path: { type: "prefix", rewrite: "/", value: "/index" },
                  headers: [{ name: "accept", type: "exact", value: "application/json" }],
                },
              },
            ],
          },
        ],
        port: 8081,
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricMeshManagementClient(credential, subscriptionId);
  const result = await client.gateway.create(
    resourceGroupName,
    gatewayResourceName,
    gatewayResourceDescription
  );
  console.log(result);
}

createOrUpdateGateway().catch(console.error);
