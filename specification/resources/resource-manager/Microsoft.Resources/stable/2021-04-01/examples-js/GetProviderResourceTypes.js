const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the resource types for a specified resource provider.
 *
 * @summary List the resource types for a specified resource provider.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/GetProviderResourceTypes.json
 */
async function getProviderResourceTypes() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceProviderNamespace = "Microsoft.TestRP";
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.providerResourceTypes.list(resourceProviderNamespace);
  console.log(result);
}

getProviderResourceTypes().catch(console.error);
