const { MongoClusterManagementClient } = require("@azure/arm-mongocluster");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing mongo cluster. The request body can contain one to many of the properties present in the normal mongo cluster definition.
 *
 * @summary updates an existing mongo cluster. The request body can contain one to many of the properties present in the normal mongo cluster definition.
 * x-ms-original-file: 2025-08-01-preview/MongoClusters_PatchCMK.json
 */
async function updatesTheCustomerManagedEncryptionKeyOnAMongoClusterResource() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MongoClusterManagementClient(credential, subscriptionId);
  const result = await client.mongoClusters.update("TestResourceGroup", "myMongoCluster", {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity2":
          {},
      },
    },
    properties: {
      encryption: {
        customerManagedKeyEncryption: {
          keyEncryptionKeyIdentity: {
            identityType: "UserAssignedIdentity",
            userAssignedIdentityResourceId:
              "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity2",
          },
          keyEncryptionKeyUrl: "https://myVault.vault.azure.net/keys/myKey2",
        },
      },
    },
  });
  console.log(result);
}
