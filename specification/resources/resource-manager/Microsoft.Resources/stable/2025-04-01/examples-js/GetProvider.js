const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the specified resource provider.
 *
 * @summary Gets the specified resource provider.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2025-04-01/examples/GetProvider.json
 */
async function getProvider() {
  const subscriptionId =
    process.env["RESOURCES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceProviderNamespace = "Microsoft.TestRP1";
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.providers.get(resourceProviderNamespace);
  console.log(result);
}
