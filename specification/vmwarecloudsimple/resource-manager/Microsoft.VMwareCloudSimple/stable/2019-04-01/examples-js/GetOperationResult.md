```javascript
const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return an async operation
 *
 * @summary Return an async operation
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetOperationResult.json
 */
async function getOperationResult() {
  const subscriptionId = "{subscription-id}";
  const regionId = "westus2";
  const referer = "https://management.azure.com/";
  const operationId = "f8e1c8f1-7d52-11e9-8e07-9a86872085ff";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.operations.get(regionId, referer, operationId);
  console.log(result);
}

getOperationResult().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-vmwarecloudsimple_3.0.0/sdk/vmwarecloudsimple/arm-vmwarecloudsimple/README.md) on how to add the SDK to your project and authenticate.
