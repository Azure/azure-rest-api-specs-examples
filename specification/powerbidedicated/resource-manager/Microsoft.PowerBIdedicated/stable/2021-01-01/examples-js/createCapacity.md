Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-powerbidedicated_3.0.1/sdk/powerbidedicated/arm-powerbidedicated/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PowerBIDedicated } = require("@azure/arm-powerbidedicated");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provisions the specified Dedicated capacity based on the configuration specified in the request.
 *
 * @summary Provisions the specified Dedicated capacity based on the configuration specified in the request.
 * x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/createCapacity.json
 */
async function createCapacity() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const resourceGroupName = "TestRG";
  const dedicatedCapacityName = "azsdktest";
  const capacityParameters = {
    administration: {
      members: ["azsdktest@microsoft.com", "azsdktest2@microsoft.com"],
    },
    location: "West US",
    sku: { name: "A1", tier: "PBIE_Azure" },
    tags: { testKey: "testValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PowerBIDedicated(credential, subscriptionId);
  const result = await client.capacities.beginCreateAndWait(
    resourceGroupName,
    dedicatedCapacityName,
    capacityParameters
  );
  console.log(result);
}

createCapacity().catch(console.error);
```
