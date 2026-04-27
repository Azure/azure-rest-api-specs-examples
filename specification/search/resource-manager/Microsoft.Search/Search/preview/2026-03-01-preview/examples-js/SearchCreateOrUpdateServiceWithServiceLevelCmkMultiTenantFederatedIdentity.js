const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a search service in the given resource group. If the search service already exists, all properties will be updated with the given values.
 *
 * @summary creates or updates a search service in the given resource group. If the search service already exists, all properties will be updated with the given values.
 * x-ms-original-file: 2026-03-01-preview/SearchCreateOrUpdateServiceWithServiceLevelCmkMultiTenantFederatedIdentity.json
 */
async function searchCreateOrUpdateServiceWithServiceLevelCmkMultiTenantFederatedIdentity() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.services.createOrUpdate("rg1", "mysearchservice", {
    location: "westus",
    tags: { "app-name": "My e-commerce app" },
    sku: { name: "standard" },
    replicaCount: 3,
    partitionCount: 1,
    hostingMode: "Default",
    computeType: "Default",
    encryptionWithCmk: {
      enforcement: "Enabled",
      serviceLevelEncryptionKey: {
        keyName: "myUserManagedEncryptionKey-createdinAzureKeyVault",
        keyVersion: "myKeyVersion-32charAlphaNumericString",
        vaultUri: "https://myKeyVault.vault.azure.net",
        identity: {
          odataType: "#Microsoft.Azure.Search.DataUserAssignedIdentity",
          userAssignedIdentity:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/user-mi",
          federatedIdentityClientId: "f83c6b1b-4d34-47e4-bb34-9d83df58b540",
        },
      },
    },
    identity: {
      type: "SystemAssigned, UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/user-mi":
          {},
      },
    },
  });
  console.log(result);
}
