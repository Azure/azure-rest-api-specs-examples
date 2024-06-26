const { ServiceFabricMeshManagementClient } = require("@azure/arm-servicefabricmesh");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a network resource with the specified name, description and properties. If a network resource with the same name exists, then it is updated with the specified description and properties.
 *
 * @summary Creates a network resource with the specified name, description and properties. If a network resource with the same name exists, then it is updated with the specified description and properties.
 * x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/networks/create_update.json
 */
async function createOrUpdateNetwork() {
  const subscriptionId =
    process.env["SERVICEFABRICMESH_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRICMESH_RESOURCE_GROUP"] || "sbz_demo";
  const networkResourceName = "sampleNetwork";
  const networkResourceDescription = {
    location: "EastUS",
    properties: {
      description: "Service Fabric Mesh sample network.",
      kind: "Local",
      networkAddressPrefix: "2.0.0.0/16",
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricMeshManagementClient(credential, subscriptionId);
  const result = await client.network.create(
    resourceGroupName,
    networkResourceName,
    networkResourceDescription
  );
  console.log(result);
}
