const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets provider operations metadata for the specified resource provider.
 *
 * @summary Gets provider operations metadata for the specified resource provider.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/GetProviderOperationsRP.json
 */
async function listProviderOperationsMetadataForResourceProvider() {
  const resourceProviderNamespace = "resourceProviderNamespace";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.providerOperationsMetadataOperations.get(resourceProviderNamespace);
  console.log(result);
}
