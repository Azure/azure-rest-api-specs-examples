Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-mariadb_2.0.1/sdk/mariadb/arm-mariadb/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new server or updates an existing server. The update action will overwrite the existing server.
 *
 * @summary Creates a new server or updates an existing server. The update action will overwrite the existing server.
 * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerCreateReplicaMode.json
 */
async function createAReplicaServer() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "TargetResourceGroup";
  const serverName = "targetserver";
  const parameters = {
    location: "westus",
    properties: {
      createMode: "Replica",
      sourceServerId:
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/MasterResourceGroup/providers/Microsoft.DBforMariaDB/servers/masterserver",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.servers.beginCreateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}

createAReplicaServer().catch(console.error);
```
