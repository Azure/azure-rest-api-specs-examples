Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-vmwarecloudsimple_3.0.0/sdk/vmwarecloudsimple/arm-vmwarecloudsimple/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of resource pools in region for private cloud
 *
 * @summary Returns list of resource pools in region for private cloud
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListResourcePools.json
 */
async function listResourcePools() {
  const subscriptionId = "{subscription-id}";
  const regionId = "westus2";
  const pcName = "myPrivateCloud";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.resourcePools.list(regionId, pcName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listResourcePools().catch(console.error);
```
