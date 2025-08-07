const { MongoClusterManagementClient } = require("@azure/arm-mongocluster");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a mongo cluster. Update overwrites all properties for the resource. To only modify some of the properties, use PATCH.
 *
 * @summary create or update a mongo cluster. Update overwrites all properties for the resource. To only modify some of the properties, use PATCH.
 * x-ms-original-file: 2025-07-01-preview/MongoClusters_CreateGeoReplica_CMK.json
 */
async function createsAReplicaMongoClusterResourceWithCustomerManagedKeyEncryptionFromASourceResource() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MongoClusterManagementClient(credential, subscriptionId);
  const result = await client.mongoClusters.createOrUpdate(
    "TestResourceGroup",
    "myReplicaMongoCluster",
    {
      identity: {
        type: "UserAssigned",
        userAssignedIdentities: {
          "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity":
            {},
        },
      },
      location: "centralus",
      properties: {
        createMode: "GeoReplica",
        replicaParameters: {
          sourceResourceId:
            "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/mySourceMongoCluster",
          sourceLocation: "eastus",
        },
        encryption: {
          customerManagedKeyEncryption: {
            keyEncryptionKeyIdentity: {
              identityType: "UserAssignedIdentity",
              userAssignedIdentityResourceId:
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity",
            },
            keyEncryptionKeyUrl: "https://myVault.vault.azure.net/keys/myKey",
          },
        },
      },
    },
  );
  console.log(result);
}
