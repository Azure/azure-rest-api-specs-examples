```javascript
const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch a subvolume
 *
 * @summary Patch a subvolume
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-01-01/examples/Subvolumes_Update.json
 */
async function subvolumesUpdate() {
  const subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = "myRG";
  const accountName = "account1";
  const poolName = "pool1";
  const volumeName = "volume1";
  const subvolumeName = "subvolume1";
  const body = { path: "/subvolumePath" };
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.subvolumes.beginUpdateAndWait(
    resourceGroupName,
    accountName,
    poolName,
    volumeName,
    subvolumeName,
    body
  );
  console.log(result);
}

subvolumesUpdate().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-netapp_16.0.0/sdk/netapp/arm-netapp/README.md) on how to add the SDK to your project and authenticate.
