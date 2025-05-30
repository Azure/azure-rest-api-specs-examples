const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates an existing (overwrite/recreate, with potential downtime) cache cluster
 *
 * @summary Creates or updates an existing (overwrite/recreate, with potential downtime) cache cluster
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2025-05-01-preview/examples/RedisEnterpriseCreate.json
 */
async function redisEnterpriseCreate() {
  const subscriptionId =
    process.env["REDISENTERPRISE_SUBSCRIPTION_ID"] || "e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f";
  const resourceGroupName = process.env["REDISENTERPRISE_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cache1";
  const parameters = {
    encryption: {
      customerManagedKeyEncryption: {
        keyEncryptionKeyIdentity: {
          identityType: "userAssignedIdentity",
          userAssignedIdentityResourceId:
            "/subscriptions/your-subscription/resourceGroups/your-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/your-identity",
        },
        keyEncryptionKeyUrl: "https://your-kv.vault.azure.net/keys/your-key/your-key-version",
      },
    },
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/yourSubscription/resourceGroups/yourRg/providers/MicrosoftManagedIdentity/userAssignedIdentities/yourIdentity":
          {},
      },
    },
    location: "West US",
    minimumTlsVersion: "1.2",
    sku: { name: "EnterpriseFlash_F300", capacity: 3 },
    tags: { tag1: "value1" },
    zones: ["1", "2", "3"],
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.redisEnterprise.beginCreateAndWait(
    resourceGroupName,
    clusterName,
    parameters,
  );
  console.log(result);
}
