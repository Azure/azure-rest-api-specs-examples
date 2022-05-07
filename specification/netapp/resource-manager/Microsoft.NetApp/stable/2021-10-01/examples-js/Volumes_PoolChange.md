Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-netapp_15.1.1/sdk/netapp/arm-netapp/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Moves volume to another pool
 *
 * @summary Moves volume to another pool
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/Volumes_PoolChange.json
 */
async function volumesAuthorizeReplication() {
  const subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = "myRG";
  const accountName = "account1";
  const poolName = "pool1";
  const volumeName = "volume1";
  const body = {
    newPoolResourceId:
      "/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.volumes.beginPoolChangeAndWait(
    resourceGroupName,
    accountName,
    poolName,
    volumeName,
    body
  );
  console.log(result);
}

volumesAuthorizeReplication().catch(console.error);
```
