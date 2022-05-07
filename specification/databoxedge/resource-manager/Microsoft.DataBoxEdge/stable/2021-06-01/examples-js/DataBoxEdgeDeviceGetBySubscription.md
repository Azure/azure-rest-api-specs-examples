Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-databoxedge_2.0.1/sdk/databoxedge/arm-databoxedge/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the Data Box Edge/Data Box Gateway devices in a subscription.
 *
 * @summary Gets all the Data Box Edge/Data Box Gateway devices in a subscription.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/DataBoxEdgeDeviceGetBySubscription.json
 */
async function dataBoxEdgeDeviceGetBySubscription() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.devices.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

dataBoxEdgeDeviceGetBySubscription().catch(console.error);
```
