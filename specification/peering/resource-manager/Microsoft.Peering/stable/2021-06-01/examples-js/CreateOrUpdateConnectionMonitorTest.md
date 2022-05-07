Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-peering_2.0.1/sdk/peering/arm-peering/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a connection monitor test with the specified name under the given subscription, resource group and peering service.
 *
 * @summary Creates or updates a connection monitor test with the specified name under the given subscription, resource group and peering service.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/CreateOrUpdateConnectionMonitorTest.json
 */
async function createOrUpdateConnectionMonitorTest() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringServiceName = "peeringServiceName";
  const connectionMonitorTestName = "connectionMonitorTestName";
  const connectionMonitorTest = {
    destination: "Example Destination",
    destinationPort: 443,
    sourceAgent: "Example Source Agent",
    testFrequencyInSec: 30,
  };
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.connectionMonitorTests.createOrUpdate(
    resourceGroupName,
    peeringServiceName,
    connectionMonitorTestName,
    connectionMonitorTest
  );
  console.log(result);
}

createOrUpdateConnectionMonitorTest().catch(console.error);
```
