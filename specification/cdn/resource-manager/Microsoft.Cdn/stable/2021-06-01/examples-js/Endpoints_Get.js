const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an existing CDN endpoint with the specified endpoint name under the specified subscription, resource group and profile.
 *
 * @summary Gets an existing CDN endpoint with the specified endpoint name under the specified subscription, resource group and profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Endpoints_Get.json
 */
async function endpointsGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.endpoints.get(resourceGroupName, profileName, endpointName);
  console.log(result);
}

endpointsGet().catch(console.error);
