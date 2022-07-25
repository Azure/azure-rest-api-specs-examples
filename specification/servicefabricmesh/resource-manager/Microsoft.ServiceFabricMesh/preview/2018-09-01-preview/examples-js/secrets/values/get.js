const { ServiceFabricMeshManagementClient } = require("@azure/arm-servicefabricmesh");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the information about the specified named secret value resources. The information does not include the actual value of the secret.
 *
 * @summary Get the information about the specified named secret value resources. The information does not include the actual value of the secret.
 * x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/secrets/values/get.json
 */
async function getSecretValue() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sbz_demo";
  const secretResourceName = "dbConnectionString";
  const secretValueResourceName = "v1";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricMeshManagementClient(credential, subscriptionId);
  const result = await client.secretValueOperations.get(
    resourceGroupName,
    secretResourceName,
    secretValueResourceName
  );
  console.log(result);
}

getSecretValue().catch(console.error);
