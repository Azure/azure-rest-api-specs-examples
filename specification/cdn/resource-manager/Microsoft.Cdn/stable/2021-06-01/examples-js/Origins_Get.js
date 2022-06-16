const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an existing origin within an endpoint.
 *
 * @summary Gets an existing origin within an endpoint.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Origins_Get.json
 */
async function originsGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const originName = "www-someDomain-net";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.origins.get(resourceGroupName, profileName, endpointName, originName);
  console.log(result);
}

originsGet().catch(console.error);
