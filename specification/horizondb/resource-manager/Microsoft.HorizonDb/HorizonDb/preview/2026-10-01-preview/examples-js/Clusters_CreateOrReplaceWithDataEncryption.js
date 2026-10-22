const { HorizonDbClient } = require("@azure/arm-horizondb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new HorizonDB cluster or replaces an existing cluster. PointInTimeRestore creates the new Cluster resource and hydrates its service-owned database storage from an available source backup. It does not clone source pools, Pool-scoped endpoints, firewall rules, parameter-group associations, or other child resources. The new Cluster receives its service-generated Cluster-scoped Endpoint named cluster. For restore, postgreSqlVersion must be omitted because the service obtains it from backup metadata; all other new-cluster configuration comes from the request or service defaults. Same-location restore requires and honors pointInTimeUtc and allows creationSourceClusterLocation to be omitted. Cross-location restore requires creationSourceClusterLocation to identify the source cluster's Azure location and requires pointInTimeUtc, but ignores the supplied timestamp and restores the latest available backup. Cross-location responses omit pointInTimeUtc.
 *
 * @summary creates a new HorizonDB cluster or replaces an existing cluster. PointInTimeRestore creates the new Cluster resource and hydrates its service-owned database storage from an available source backup. It does not clone source pools, Pool-scoped endpoints, firewall rules, parameter-group associations, or other child resources. The new Cluster receives its service-generated Cluster-scoped Endpoint named cluster. For restore, postgreSqlVersion must be omitted because the service obtains it from backup metadata; all other new-cluster configuration comes from the request or service defaults. Same-location restore requires and honors pointInTimeUtc and allows creationSourceClusterLocation to be omitted. Cross-location restore requires creationSourceClusterLocation to identify the source cluster's Azure location and requires pointInTimeUtc, but ignores the supplied timestamp and restores the latest available backup. Cross-location responses omit pointInTimeUtc.
 * x-ms-original-file: 2026-10-01-preview/Clusters_CreateOrReplaceWithDataEncryption.json
 */
async function createAHorizonDBClusterWithCustomerManagedKeyDataEncryption() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new HorizonDbClient(credential, subscriptionId);
  const result = await client.clusters.createOrReplace(
    "example-resource-group",
    "example-cluster",
    {
      location: "eastus2",
      identity: {
        type: "UserAssigned",
        userAssignedIdentities: {
          "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/identity-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/primary-encryption-identity":
            {},
        },
      },
      properties: {
        postgreSqlVersion: "17",
        backup: { geoRedundancy: "Disabled", retentionDays: 14 },
        dataEncryption: {
          type: "AzureKeyVault",
          primaryKeyUri: "https://primary-key-vault.vault.azure.net/keys/horizondb-primary-key",
          primaryUserAssignedIdentityId:
            "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/identity-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/primary-encryption-identity",
        },
      },
    },
  );
  console.log(result);
}
