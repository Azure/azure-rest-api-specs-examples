Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-powerbidedicated_3.0.1/sdk/powerbidedicated/arm-powerbidedicated/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PowerBIDedicated } = require("@azure/arm-powerbidedicated");
const { DefaultAzureCredential } = require("@azure/identity");

async function getDetailsOfAnAutoScaleVCore() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const resourceGroupName = "TestRG";
  const vcoreName = "testvcore";
  const credential = new DefaultAzureCredential();
  const client = new PowerBIDedicated(credential, subscriptionId);
  const result = await client.autoScaleVCores.get(resourceGroupName, vcoreName);
  console.log(result);
}

getDetailsOfAnAutoScaleVCore().catch(console.error);
```
