const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an existing origin group within an endpoint.
 *
 * @summary Gets an existing origin group within an endpoint.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/OriginGroups_Get.json
 */
async function originGroupsGet() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const originGroupName = "originGroup1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.originGroups.get(
    resourceGroupName,
    profileName,
    endpointName,
    originGroupName
  );
  console.log(result);
}
