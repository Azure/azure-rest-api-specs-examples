```javascript
const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of certificates.
 *
 * @summary Returns the list of certificates.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_listcertificates.json
 */
async function certificatesListByIotHub() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "testhub";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.certificates.listByIotHub(resourceGroupName, resourceName);
  console.log(result);
}

certificatesListByIotHub().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-iothub_6.1.2/sdk/iothub/arm-iothub/README.md) on how to add the SDK to your project and authenticate.
