```javascript
const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Cache.
 *
 * @summary Create or update a Cache.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/Caches_CreateOrUpdate.json
 */
async function cachesCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "scgroup";
  const cacheName = "sc1";
  const cache = {
    cacheSizeGB: 3072,
    directoryServicesSettings: {
      activeDirectory: {
        cacheNetBiosName: "contosoSmb",
        credentials: { password: "<password>", username: "consotoAdmin" },
        domainName: "contosoAd.contoso.local",
        domainNetBiosName: "contosoAd",
        primaryDnsIpAddress: "192.0.2.10",
        secondaryDnsIpAddress: "192.0.2.11",
      },
      usernameDownload: {
        credentials: {
          bindDn: "cn=ldapadmin,dc=contosoad,dc=contoso,dc=local",
          bindPassword: "<bindPassword>",
        },
        extendedGroups: true,
        ldapBaseDN: "dc=contosoad,dc=contoso,dc=local",
        ldapServer: "192.0.2.12",
        usernameSource: "LDAP",
      },
    },
    encryptionSettings: {
      keyEncryptionKey: {
        keyUrl: "https://keyvault-cmk.vault.azure.net/keys/key2047/test",
        sourceVault: {
          id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.KeyVault/vaults/keyvault-cmk",
        },
      },
    },
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000000000000000000000000000/resourceGroups/scgroup/providers/MicrosoftManagedIdentity/userAssignedIdentities/identity1":
          {},
      },
    },
    location: "westus",
    securitySettings: {
      accessPolicies: [
        {
          name: "default",
          accessRules: [
            {
              access: "rw",
              rootSquash: false,
              scope: "default",
              submountAccess: true,
              suid: false,
            },
          ],
        },
      ],
    },
    sku: { name: "Standard_2G" },
    subnet:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/sub1",
    tags: { dept: "Contoso" },
  };
  const options = { cache };
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const result = await client.caches.beginCreateOrUpdateAndWait(
    resourceGroupName,
    cacheName,
    options
  );
  console.log(result);
}

cachesCreateOrUpdate().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storagecache_5.1.0/sdk/storagecache/arm-storagecache/README.md) on how to add the SDK to your project and authenticate.
