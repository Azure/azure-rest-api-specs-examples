const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Commit the migrated Azure Frontdoor(Standard/Premium) profile.
 *
 * @summary Commit the migrated Azure Frontdoor(Standard/Premium) profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_MigrationCommit.json
 */
async function profilesMigrationCommit() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.profiles.beginMigrationCommitAndWait(resourceGroupName, profileName);
  console.log(result);
}
