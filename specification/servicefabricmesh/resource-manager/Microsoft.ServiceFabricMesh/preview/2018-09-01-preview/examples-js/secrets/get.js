const { ServiceFabricMeshManagementClient } = require("@azure/arm-servicefabricmesh");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the information about the secret resource with the given name. The information include the description and other properties of the secret.
 *
 * @summary Gets the information about the secret resource with the given name. The information include the description and other properties of the secret.
 * x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/secrets/get.json
 */
async function getSecret() {
  const subscriptionId =
    process.env["SERVICEFABRICMESH_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRICMESH_RESOURCE_GROUP"] || "sbz_demo";
  const secretResourceName = "dbConnectionString";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricMeshManagementClient(credential, subscriptionId);
  const result = await client.secret.get(resourceGroupName, secretResourceName);
  console.log(result);
}
