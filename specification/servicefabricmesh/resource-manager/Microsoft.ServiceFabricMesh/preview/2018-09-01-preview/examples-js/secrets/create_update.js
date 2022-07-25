const { ServiceFabricMeshManagementClient } = require("@azure/arm-servicefabricmesh");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a secret resource with the specified name, description and properties. If a secret resource with the same name exists, then it is updated with the specified description and properties.
 *
 * @summary Creates a secret resource with the specified name, description and properties. If a secret resource with the same name exists, then it is updated with the specified description and properties.
 * x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/secrets/create_update.json
 */
async function createOrUpdateSecret() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sbz_demo";
  const secretResourceName = "dbConnectionString";
  const secretResourceDescription = {
    location: "EastUS",
    properties: {
      description: "Mongo DB connection string for backend database!",
      contentType: "text/plain",
      kind: "inlinedValue",
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricMeshManagementClient(credential, subscriptionId);
  const result = await client.secret.create(
    resourceGroupName,
    secretResourceName,
    secretResourceDescription
  );
  console.log(result);
}

createOrUpdateSecret().catch(console.error);
