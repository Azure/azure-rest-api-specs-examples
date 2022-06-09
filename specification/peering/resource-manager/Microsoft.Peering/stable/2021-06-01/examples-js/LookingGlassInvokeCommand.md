```javascript
const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Run looking glass functionality
 *
 * @summary Run looking glass functionality
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/LookingGlassInvokeCommand.json
 */
async function callLookingGlassToExecuteACommand() {
  const subscriptionId = "subId";
  const command = "Traceroute";
  const sourceType = "AzureRegion";
  const sourceLocation = "West US";
  const destinationIP = "0.0.0.0";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.lookingGlass.invoke(
    command,
    sourceType,
    sourceLocation,
    destinationIP
  );
  console.log(result);
}

callLookingGlassToExecuteACommand().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-peering_2.0.1/sdk/peering/arm-peering/README.md) on how to add the SDK to your project and authenticate.
