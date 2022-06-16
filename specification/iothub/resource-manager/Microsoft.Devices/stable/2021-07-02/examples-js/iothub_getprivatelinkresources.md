```javascript
const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified private link resource for the given IotHub
 *
 * @summary Get the specified private link resource for the given IotHub
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_getprivatelinkresources.json
 */
async function privateLinkResourcesList() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "testHub";
  const groupId = "iotHub";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.privateLinkResourcesOperations.get(
    resourceGroupName,
    resourceName,
    groupId
  );
  console.log(result);
}

privateLinkResourcesList().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-iothub_6.1.2/sdk/iothub/arm-iothub/README.md) on how to add the SDK to your project and authenticate.
