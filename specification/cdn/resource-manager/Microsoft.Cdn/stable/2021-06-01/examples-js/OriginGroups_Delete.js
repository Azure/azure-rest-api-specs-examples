const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing origin group within an endpoint.
 *
 * @summary Deletes an existing origin group within an endpoint.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/OriginGroups_Delete.json
 */
async function originGroupsDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const originGroupName = "originGroup1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.originGroups.beginDeleteAndWait(
    resourceGroupName,
    profileName,
    endpointName,
    originGroupName
  );
  console.log(result);
}

originGroupsDelete().catch(console.error);
