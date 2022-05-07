Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storsimple8000series_2.0.1/sdk/storsimple8000series/arm-storsimple8000series/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the manager.
 *
 * @summary Creates or updates the manager.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/ManagersCreateOrUpdate.json
 */
async function managersCreateOrUpdate() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest2";
  const parameters = {
    cisIntrinsicSettings: { type: "GardaV1" },
    location: "westus",
    sku: { name: "Standard" },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.managers.createOrUpdate(resourceGroupName, managerName, parameters);
  console.log(result);
}

managersCreateOrUpdate().catch(console.error);
```
