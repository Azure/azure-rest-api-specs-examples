const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the provider permissions.
 *
 * @summary Get the provider permissions.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/GetProviderPermissions.json
 */
async function getProviderResourceTypes() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceProviderNamespace = "Microsoft.TestRP";
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.providers.providerPermissions(resourceProviderNamespace);
  console.log(result);
}

getProviderResourceTypes().catch(console.error);
