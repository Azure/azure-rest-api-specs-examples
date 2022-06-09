```javascript
const { PowerBIDedicated } = require("@azure/arm-powerbidedicated");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the Dedicated capacities for the given subscription.
 *
 * @summary Lists all the Dedicated capacities for the given subscription.
 * x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/listCapacitiesInSubscription.json
 */
async function getDetailsOfACapacity() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const credential = new DefaultAzureCredential();
  const client = new PowerBIDedicated(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.capacities.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getDetailsOfACapacity().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-powerbidedicated_3.0.1/sdk/powerbidedicated/arm-powerbidedicated/README.md) on how to add the SDK to your project and authenticate.
