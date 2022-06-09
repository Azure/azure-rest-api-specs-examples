```javascript
const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a Cache instance.
 *
 * @summary Update a Cache instance.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/Caches_Update.json
 */
async function cachesUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "scgroup";
  const cacheName = "sc1";
  const cache = {
    cacheSizeGB: 3072,
    directoryServicesSettings: {
      activeDirectory: {
        cacheNetBiosName: "contosoSmb",
        domainName: "contosoAd.contoso.local",
        domainNetBiosName: "contosoAd",
        primaryDnsIpAddress: "192.0.2.10",
        secondaryDnsIpAddress: "192.0.2.11",
      },
      usernameDownload: { extendedGroups: true, usernameSource: "AD" },
    },
    location: "westus",
    networkSettings: {
      dnsSearchDomain: "contoso.com",
      dnsServers: ["10.1.22.33", "10.1.12.33"],
      mtu: 1500,
      ntpServer: "time.contoso.com",
    },
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
        {
          name: "restrictive",
          accessRules: [
            {
              access: "rw",
              filter: "10.99.3.145",
              rootSquash: false,
              scope: "host",
              submountAccess: true,
              suid: true,
            },
            {
              access: "rw",
              filter: "10.99.1.0/24",
              rootSquash: false,
              scope: "network",
              submountAccess: true,
              suid: true,
            },
            {
              access: "no",
              anonymousGID: "65534",
              anonymousUID: "65534",
              rootSquash: true,
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
  const result = await client.caches.update(resourceGroupName, cacheName, options);
  console.log(result);
}

cachesUpdate().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storagecache_5.1.0/sdk/storagecache/arm-storagecache/README.md) on how to add the SDK to your project and authenticate.
