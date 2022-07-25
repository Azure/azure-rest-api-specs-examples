const { ServiceFabricMeshManagementClient } = require("@azure/arm-servicefabricmesh");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about all secret value resources of the specified secret resource. The information includes the names of the secret value resources, but not the actual values.
 *
 * @summary Gets information about all secret value resources of the specified secret resource. The information includes the names of the secret value resources, but not the actual values.
 * x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/secrets/values/list.json
 */
async function listSecretValues() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sbz_demo";
  const secretResourceName = "dbConnectionString";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricMeshManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.secretValueOperations.list(resourceGroupName, secretResourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSecretValues().catch(console.error);
