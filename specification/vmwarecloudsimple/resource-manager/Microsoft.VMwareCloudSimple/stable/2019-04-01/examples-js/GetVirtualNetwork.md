```javascript
const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return virtual network by its name
 *
 * @summary Return virtual network by its name
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetVirtualNetwork.json
 */
async function getVirtualNetwork() {
  const subscriptionId = "{subscription-id}";
  const regionId = "westus2";
  const pcName = "myPrivateCloud";
  const virtualNetworkName = "dvportgroup-19";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.virtualNetworks.get(regionId, pcName, virtualNetworkName);
  console.log(result);
}

getVirtualNetwork().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-vmwarecloudsimple_3.0.0/sdk/vmwarecloudsimple/arm-vmwarecloudsimple/README.md) on how to add the SDK to your project and authenticate.
