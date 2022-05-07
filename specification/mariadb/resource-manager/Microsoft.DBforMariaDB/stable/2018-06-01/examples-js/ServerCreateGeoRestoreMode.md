Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-mariadb_2.0.1/sdk/mariadb/arm-mariadb/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new server or updates an existing server. The update action will overwrite the existing server.
 *
 * @summary Creates a new server or updates an existing server. The update action will overwrite the existing server.
 * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerCreateGeoRestoreMode.json
 */
async function createAServerAsAGeoRestore() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "TargetResourceGroup";
  const serverName = "targetserver";
  const parameters = {
    location: "westus",
    properties: {
      createMode: "GeoRestore",
      sourceServerId:
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMariaDB/servers/sourceserver",
    },
    sku: {
      name: "GP_Gen5_2",
      capacity: 2,
      family: "Gen5",
      tier: "GeneralPurpose",
    },
    tags: { elasticServer: "1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.servers.beginCreateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}

createAServerAsAGeoRestore().catch(console.error);
```
