const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing origin group within a profile.
 *
 * @summary Deletes an existing origin group within a profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOriginGroups_Delete.json
 */
async function afdOriginGroupsDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const originGroupName = "origingroup1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdOriginGroups.beginDeleteAndWait(
    resourceGroupName,
    profileName,
    originGroupName
  );
  console.log(result);
}

afdOriginGroupsDelete().catch(console.error);
