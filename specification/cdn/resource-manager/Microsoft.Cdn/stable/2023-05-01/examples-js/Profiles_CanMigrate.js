const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks if CDN profile can be migrated to Azure Frontdoor(Standard/Premium) profile.
 *
 * @summary Checks if CDN profile can be migrated to Azure Frontdoor(Standard/Premium) profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Profiles_CanMigrate.json
 */
async function profilesCanMigrate() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const canMigrateParameters = {
    classicResourceReference: {
      id: "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Network/frontdoors/frontdoorname",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.profiles.beginCanMigrateAndWait(
    resourceGroupName,
    canMigrateParameters
  );
  console.log(result);
}
