```javascript
const { PowerBIDedicated } = require("@azure/arm-powerbidedicated");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the current state of the specified auto scale v-core.
 *
 * @summary Updates the current state of the specified auto scale v-core.
 * x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/updateAutoScaleVCore.json
 */
async function updateAutoScaleVCoreParameters() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const resourceGroupName = "TestRG";
  const vcoreName = "testvcore";
  const vCoreUpdateParameters = {
    capacityLimit: 20,
    sku: { name: "AutoScale", capacity: 0, tier: "AutoScale" },
    tags: { testKey: "testValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PowerBIDedicated(credential, subscriptionId);
  const result = await client.autoScaleVCores.update(
    resourceGroupName,
    vcoreName,
    vCoreUpdateParameters
  );
  console.log(result);
}

updateAutoScaleVCoreParameters().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-powerbidedicated_3.0.1/sdk/powerbidedicated/arm-powerbidedicated/README.md) on how to add the SDK to your project and authenticate.
