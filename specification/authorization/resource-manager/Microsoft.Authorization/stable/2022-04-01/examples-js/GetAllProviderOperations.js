const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets provider operations metadata for all resource providers.
 *
 * @summary Gets provider operations metadata for all resource providers.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/GetAllProviderOperations.json
 */
async function listProviderOperationsMetadataForAllResourceProviders() {
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.providerOperationsMetadataOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
